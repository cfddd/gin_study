package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(logrus.GetLevel())
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Trace("trace msg") //不输出
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg") //在这里已经退出
	logrus.Panic("panic msg") //不输出
}
