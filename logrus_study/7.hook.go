package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type MyHook struct {
}

func (MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.PanicLevel}
}

// Fire 只在logrus.ErrorLevel, logrus.PanicLevel等级调用以下内容，输出到文件中
func (MyHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = "CFD"
	//fmt.Println(entry)

	file, _ := os.OpenFile("logrus_study/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	line, _ := entry.String()
	file.Write([]byte(line))

	return nil
}

func main() {
	logrus.AddHook(&MyHook{})

	logrus.Warnln("hello")
	logrus.Errorln("hello")
}
