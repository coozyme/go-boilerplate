package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logging struct {
	Log      *logrus.Logger
	FileName string
}

func NewLogger(log *logrus.Logger, filename string) *Logging {
	return &Logging{Log: log, FileName: filename}
}

func (l *Logging) InfoLog(message ...interface{}) {
	log := l.Log
	file, _ := os.OpenFile(l.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Info(message)
}

func (l *Logging) ErrorLog(message ...interface{}) {
	log := l.Log
	file, _ := os.OpenFile(l.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Error(message)
}

func (l *Logging) FatalLog(message ...interface{}) {
	log := l.Log
	file, _ := os.OpenFile(l.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Fatal(message)
}
