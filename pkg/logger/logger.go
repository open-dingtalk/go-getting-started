package logger

/**
 * @Author linya.jj
 * @Date 2022/10/23 2:16 PM
 */

import (
	"fmt"
	"time"
)

type ILogger interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type ConsoleLogger struct {
}

func (l *ConsoleLogger) Info(format string, args ...interface{}) {
	fmt.Println(time.Now().String(), "[INFO]", fmt.Sprintf(format, args...))
}

func (l *ConsoleLogger) Warn(format string, args ...interface{}) {
	fmt.Println(time.Now().String(), "[INFO]", fmt.Sprintf(format, args...))
}

func (l *ConsoleLogger) Error(format string, args ...interface{}) {
	fmt.Println(time.Now().String(), "[ERROR]", time.Now().String(), fmt.Sprintf(format, args...))
}

var (
	defaultLogger ILogger
)

func ActiveInitLogger() {
	defaultLogger = &ConsoleLogger{}
}

func GetDefaultLogger() ILogger {
	if defaultLogger != nil {
		return defaultLogger
	}

	defaultLogger = &ConsoleLogger{}
	return defaultLogger
}
