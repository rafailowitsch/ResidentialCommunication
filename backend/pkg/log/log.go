// log/logger.go
package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)

	logrus.SetFormatter(logger.Formatter)
	logrus.SetOutput(logger.Out)
	logrus.SetLevel(logger.Level)

	logrus.WithFields(logrus.Fields{
		"module":   "log",
		"function": "init",
	}).Info("Logger initialized")
}

func LogInfo(operation, message string, fields logrus.Fields) {
	fields["operation"] = operation
	logrus.WithFields(fields).Info(message)
}

func LogError(operation, message string, err error, fields logrus.Fields) {
	fields["operation"] = operation
	fields["error"] = err
	logrus.WithFields(fields).Error(message)
}
