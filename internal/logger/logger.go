package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New(logLevel string) *Logger {
	log := logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Fatal(err)
	}
	log.Level = level

	log.Out = os.Stdout

	return &Logger{
		logger: log,
	}
}

func (l *Logger) getEntry() *logrus.Entry {
	return l.logger.WithFields(logrus.Fields{})
}

func (l *Logger) Error(message string) {
	l.getEntry().Error(message)
}

func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

func (l *Logger) Warning(message string) {
	l.logger.Warning(message)
}

func (l *Logger) Debug(message string) {
	l.logger.Debug(message)
}

func (l *Logger) Fatal(message string) {
	l.logger.Fatal(message)
}
