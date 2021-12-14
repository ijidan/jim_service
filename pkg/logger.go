package pkg

import (
	"github.com/sirupsen/logrus"
	"io"
	"jim/config"
	"os"
	"sync"
)

var (
	onceLogger     sync.Once
	instanceLogger *logrus.Logger
)

func GetLoggerInstance(config *config.Config, root string) *logrus.Logger {
	onceLogger.Do(func() {
		logfile := root + "/" + config.Http.Log
		writer1 := os.Stdout
		writer2, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		instanceLogger = logrus.New()
		instanceLogger.SetFormatter(&logrus.JSONFormatter{})
		instanceLogger.SetOutput(io.MultiWriter(writer1, writer2))
	})
	return instanceLogger
}
