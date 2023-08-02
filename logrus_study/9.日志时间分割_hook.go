package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type _FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
}

func (hook _FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook _FileDateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02_15-04")

	line, _ := entry.String()

	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	//时间不等,重新创建
	hook.file.Close()
	//InitFile("logrus_study/logHook", "CFD")//可以吧？
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)
	filename := fmt.Sprintf("%s/%s/%s-%s.log", hook.logPath, timer, hook.appName, time.Now().Format("2006-01-02_15-04"))

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

func InitFile(logPath, appName string) {
	fileDate := time.Now().Format("2006-01-02_15-04")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}

	//创建目录下文件
	filename := fmt.Sprintf("%s/%s/%s-%s.log", logPath, fileDate, appName, time.Now().Format("2006-01-02_15-04"))
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}

	//用hook的方式调用，hook就是一个附带成员和方法的东西，这样输出会在logrus里面会自动用到
	filehook := _FileDateHook{file, logPath, fileDate, appName}
	logrus.AddHook(&filehook)
}

func main() {
	InitFile("logrus_study/logHook", "CFD")

	for {
		logrus.Warnln("hello")
		time.Sleep(5 * time.Second)
		logrus.Errorln("hello")
	}

	logrus.Warnln("hello")
	logrus.Errorln("hello")
}
