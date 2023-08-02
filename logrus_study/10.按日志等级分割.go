/*
只分 err war info
*/
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type FileLevelHook struct {
	file     *os.File
	errFile  *os.File
	warFile  *os.File
	infoFile *os.File
}

func (hook FileLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook FileLevelHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	hook.file.Write([]byte(line))
	return nil
}

func InitLevel(logPath string) {
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

}
