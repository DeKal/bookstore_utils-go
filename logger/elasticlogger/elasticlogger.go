package elasticlogger

import (
	"fmt"

	"github.com/DeKal/bookstore_utils-go/logger"
)

const (
	// LoggerInfo logger type Info
	LoggerInfo = "INFO"
	// LoggerError logger type Error
	LoggerError = "ERROR"
)

// LoggerInterface define logger interface for elastic search
type LoggerInterface interface {
	Printf(format string, v ...interface{})
}

type esInfoLogger struct {
}

func (*esInfoLogger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		logger.Info(format)
	}
	logger.Info(fmt.Sprintf(format, v...))
}

type esErrorLogger struct {
}

func (*esErrorLogger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		logger.Error(format, nil)
	}
	logger.Error(fmt.Sprintf(format, v...), nil)
}

// NewLogger create new Eslogger
func NewLogger(loggerType string) LoggerInterface {
	switch loggerType {
	case LoggerError:
		return &esErrorLogger{}
	case LoggerInfo:
		return &esInfoLogger{}
	}
	return &esInfoLogger{}
}
