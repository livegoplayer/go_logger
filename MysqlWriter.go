package logger

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	myHelper "github.com/livegoplayer/go_helper"
)

//手动实现一个mysql_writer,作为输出流对象传递到log
type MysqlWriter struct {
	DbConnection *sql.DB
	host         string
	port         string
	dbName       string
	tableName    string
}

//设计mongo日志表存储格式
func GetMysqlWriter(host string, port string, dbName string, tableName string, username string, password string) *MysqlWriter {
	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "3306"
	}

	if tableName == "" {
		tableName = "us_app_log"
	}

	if username == "" {
		username = "root"
	}

	mw := &MysqlWriter{
		host:      host,
		port:      port,
		dbName:    dbName,
		tableName: tableName,
	}

	var err error
	mw.DbConnection, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		fmt.Printf("err:" + err.Error())
	}
	//设置数据库最大连接数
	mw.DbConnection.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	mw.DbConnection.SetMaxIdleConns(10)
	//验证连接
	if err := mw.DbConnection.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connect success")

	return mw
}

func (mw *MysqlWriter) Write(p []byte) (n int, err error) {
	n = 0
	err = nil

	//解析出level_no
	levelNo := myHelper.GetJsonVal(p, "levelNo")

	//开启事务
	tx, err := mw.DbConnection.Begin()
	if err != nil {
		fmt.Println("获取事务失败")
		return
	}

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO `" + mw.tableName + "` (`message`, `level`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare failed")
		return
	}

	//将参数传递到sql语句中并且执行
	_, err = stmt.Exec(p, levelNo)
	if err != nil {
		fmt.Println("Exec failed")
		return
	}

	//将事务提交
	err = tx.Commit()
	//获得上一个插入自增的id
	if err != nil {
		fmt.Println("commit failed")
		return
	}

	return len(p), nil
}
