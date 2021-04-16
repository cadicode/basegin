package base

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

var Logger ILogger

// ILogger define the log actions
type ILogger interface {
	// WriteError is for error record
	WriteError(err interface{}, trace string, param interface{})

	// WriteInfo is for temp debug info
	WriteInfo(param interface{})
}

// SetLogger is set for logger instance
func SetLogger(log ILogger) {
	if log != nil {
		Logger = log
	}
}

type LogrusLogger struct {
}

// NewLogrusLogger is init func for an new instance
func NewLogrusLogger(logPath string) (*LogrusLogger, error) {
	if len(logPath) == 0 {
		return nil, errors.New("logPath is nil")
	}
	fileHook := NewFileHook(logPath)
	logrus.AddHook(fileHook)
	return &LogrusLogger{}, nil
}

// WriteError for error info
func (logger *LogrusLogger) WriteError(err interface{}, trace string, param interface{}) {

	entry := logrus.WithField("Trace", trace)
	if param != nil {
		entry = entry.WithField("AdditionalInfo", fmt.Sprintf("%+v", param))
	}
	entry.Error(err)
}

// WriteInfo for temp debug info
func (logger *LogrusLogger) WriteInfo(param interface{}) {
	logrus.Info(param)
}
