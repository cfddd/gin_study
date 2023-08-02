package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

const (
	ccBlock  = 0
	ccRed    = 1
	ccGreed  = 2
	ccYellow = 3
	ccBlue   = 4
	ccPurple = 5
	ccCyan   = 6
	ccGrey   = 7
)

type MyFormatter struct {
	Prefix string
}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var color int
	switch entry.Level {
	case logrus.ErrorLevel:
		color = ccRed
	case logrus.WarnLevel:
		color = ccYellow
	case logrus.InfoLevel:
		color = ccBlue
	case logrus.DebugLevel:
		color = ccCyan
	case logrus.TraceLevel:
		color = ccGrey
	default:
		color = ccGrey
	}
	//文件输入输出缓冲区
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	//时间格式化
	formatTime := entry.Time.Format("2006-01-02 15:04:05")
	//自定义文件路径
	funcVal := entry.Caller.Function

	//文件行号
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	//自定义文件格式
	fmt.Fprintf(b, "[%s] \033[3%dm[%s]\033[0m [%s] [%s] [%s] [%s]\n", f.Prefix, color, entry.Level, formatTime, fileVal, funcVal, entry.Message)

	return b.Bytes(), nil
}

func main() {
	//设置行号
	logrus.SetReportCaller(true)
	//传递格式和参数相关
	logrus.SetFormatter(&MyFormatter{Prefix: "CFD"})

	logrus.SetLevel(logrus.TraceLevel)
	logrus.Infoln("hello")
	logrus.Errorln("hello")
	logrus.Warnln("hello")
	logrus.Debugln("hello")
	logrus.Traceln("hello")

}
