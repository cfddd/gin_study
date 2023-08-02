package main

import "github.com/sirupsen/logrus"

func main() {

	//两种方法使用自己的字段
	log := logrus.WithField("app", "study").WithField("servic", "logrus")
	log = log.WithFields(logrus.Fields{
		"user": "21",
		"ip":   "192.168.200.254",
	})

	log.Errorf("你好")

}
