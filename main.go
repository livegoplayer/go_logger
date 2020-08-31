package main

import (
	"github.com/livegoplayer/go_logger/logger"
)

func main() {
	testRabbitmq()
}

func testRabbitmq() {
	rm := logger.GetRabbitMqLogger("amqp://guest:guest@139.224.132.234:5670/", "log_go_user", "log", logger.GO_USER)
	logger.SetLogger(rm)

	logger.Info("1")
}
