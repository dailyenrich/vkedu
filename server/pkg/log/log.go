package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"vke/util"
)

var (
	log     = logrus.New()
	logPath = "runtime/log/"
	fileName = "%s.log"
	fileNameFormat = "20060102"
)

func Logger() *logrus.Logger {
	return log
}

func setup()  {
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		PrettyPrint:       true,
	})

	log.SetReportCaller(true)

	setLogfile()
}

func setLogfile()  {
	logFile := fmt.Sprintf(logPath+fileName, time.Now().Format(fileNameFormat))
	if !util.FileExist(logPath) {
		if err := os.MkdirAll(logPath, 0666); err != nil {
			panic(err)
		}
	}
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)
}

func Init()  {
	setup()
}

func Info(args ...interface{})  {
	log.Infoln(args...)
}

func Debug(args ...interface{})  {
	log.Debugln(args...)
}

func Warn(args ...interface{})  {
	log.Warnln(args...)
}
