package writer

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

//文件log单例
//完全日志输出的模式
func GetFileWriter(path string, level logrus.Level, cleanTime time.Duration, splitTime time.Duration) *rotatelogs.RotateLogs {
	if path == "" {
		path = "../logs"
	}

	if !Exists(filepath.FromSlash(path)) {
		err := os.MkdirAll(path, os.ModeDir)
		_ = os.Chmod(path, os.ModePerm)
		if err != nil {
			panic("创建日志文件目录失败")
		}
	}

	levelName, err := level.MarshalText()
	if err != nil {
		panic(err)
	}

	resPath := filepath.Join(path, string(levelName)+"_%Y%m%d%H%M.log")
	linkPath := filepath.Join(path, string(levelName)+".log")

	//设置对某个地方的日志进行分割
	writer, err := rotatelogs.New(
		filepath.Join(resPath),
		//每次从这个位置清除分离日志
		rotatelogs.WithLinkName(linkPath),
		//每30天清除一次日志
		rotatelogs.WithMaxAge(cleanTime),
		//每一天分离一次日志文件
		rotatelogs.WithRotationTime(splitTime),
	)

	if err != nil {
		//输出到控制台
		_, _ = fmt.Fprint(os.Stdin, "error:"+err.Error())
	}

	return writer
}
