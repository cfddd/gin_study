package main

import "github.com/sirupsen/logrus"

func main() {

	//设置提示格式，默认TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("trace msg") //不输出
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg") //在这里已经退出
	logrus.Panic("panic msg") //不输出

}
