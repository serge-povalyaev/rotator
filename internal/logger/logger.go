package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New(logLevel, filePath string) *Logger {
	log := logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Fatal(err)
	}
	log.Level = level

	log.Out = os.Stdout
	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0o666)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Ошибка в работе с файлом")
		}
	}

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
