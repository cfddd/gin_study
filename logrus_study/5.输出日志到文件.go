package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	//加
	file, _ := os.OpenFile("logrus_study/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	//同时输出日志和文件
	logrus.SetOutput(io.MultiWriter(file,
		os.Stdout))

	logrus.Infof("hello")
	logrus.Error("hello")
	logrus.Error("k")
	logrus.Error("jk")
}
