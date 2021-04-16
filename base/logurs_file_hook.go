package base

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// FileHook is a hook for writing log in local file
type FileHook struct {
	LogFolderPath string
}

var (
	logFileDay   = 0
	errorLogPath string
	debugLogPath string
)

// NewFileHook is an new method for new a file hook
func NewFileHook(logFolderPath string) *FileHook {
	logFolderPath = strings.Trim(logFolderPath, " ")
	if logFolderPath == "" {
		panic("logFolderPath is null")
	}
	return &FileHook{
		LogFolderPath: logFolderPath,
	}
}

// Levels is the method must defined in hook
func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire is the method must defined in hook
func (hook *FileHook) Fire(entry *logrus.Entry) error {
	logFilePath := hook.GetLogFilePath(entry.Level)
	fileObj, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer fileObj.Close()

	info, err := entry.String()
	if err != nil {
		fileObj.WriteString(err.Error() + "\n")
		fileObj.WriteString(fmt.Sprintf("trace:%+v", err) + "\n")
	} else {
		fileObj.WriteString(info + "\n")
	}
	return nil
}

// GetLogFilePath is a func to get log file path
func (hook *FileHook) GetLogFilePath(level logrus.Level) string {
	if logFileDay == 0 || logFileDay != time.Now().Day() { //first time to get path
		defer func() {
			logFileDay = time.Now().Day()
		}()
		errorLogPath = fmt.Sprintf("%s%s", hook.LogFolderPath, getLogName(logrus.ErrorLevel))
		debugLogPath = fmt.Sprintf("%s%s", hook.LogFolderPath, getLogName(logrus.DebugLevel))
	}
	switch level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		return errorLogPath
	default:
		return debugLogPath
	}
}

func getLogName(level logrus.Level) string {

	return fmt.Sprintf("%s.%s.log", time.Now().Format("20060102"), getLogFileExtension(level))
}

func getLogFileExtension(level logrus.Level) string {
	switch level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		return "error"
	default:
		return "debug"

	}
}
