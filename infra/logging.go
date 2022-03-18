package infra 

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logging struct {
	*logrus.Logger
}

var loggingSingleton *Logging

func InitLogging() (*Logging, error) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{})

	log.SetReportCaller(true)

	log.SetLevel(logrus.WarnLevel)

	log.Out = os.Stdout

	file, err := os.OpenFile("log/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	loggingSingleton = &Logging{log}	

	return loggingSingleton, nil
}

func GetLogging() *Logging {
	if loggingSingleton == nil {
		panic("can not init logging")
	}
	return loggingSingleton
}