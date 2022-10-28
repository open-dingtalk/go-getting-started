package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/23 2:16 PM
 */

func TestConsoleLogger_Info(t *testing.T) {
	cl := &ConsoleLogger{}
	cl.Info("info test[%s]", "001")
}

func TestConsoleLogger_Warn(t *testing.T) {
	cl := &ConsoleLogger{}
	cl.Warn("info test[%s]", "001")
}

func TestConsoleLogger_Error(t *testing.T) {
	cl := &ConsoleLogger{}
	cl.Error("info test[%s]", "001")
}

func TestActiveInitLogger(t *testing.T) {
	ActiveInitLogger()
	assert.NotNil(t, GetDefaultLogger())
}

func TestGetDefaultLogger_IfExist(t *testing.T) {
	defaultLogger = &ConsoleLogger{}
	assert.NotNil(t, GetDefaultLogger())
}

func TestGetDefaultLogger_IfNotExist(t *testing.T) {
	defaultLogger = nil
	assert.NotNil(t, GetDefaultLogger())
}
