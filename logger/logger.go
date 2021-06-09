// Package logger contains some simple logging helpers
package logger

import (
	"github.com/sirupsen/logrus"
)

// New initializes a new logger
// with JSON-formatted output
func New(level logrus.Level) *logrus.Logger {
	return &logrus.Logger{
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}
}
