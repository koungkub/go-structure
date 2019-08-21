package connection

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetLogConnection() *logrus.Entry {

	l := logrus.New()

	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&logrus.JSONFormatter{})

	log := l.WithFields(logrus.Fields{
		"service": viper.GetString("APP.NAME"),
	})

	return log
}
