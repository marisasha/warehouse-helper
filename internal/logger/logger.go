package logger

import "github.com/sirupsen/logrus"

func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}
