package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLoggerConstants(t *testing.T) {
	assert.EqualValues(t, "LOG_LEVEL", envLogLevel)
	assert.EqualValues(t, "LOG_OUTPUT", envLogOutput)
}

func TestGetLevelDebug(t *testing.T) {
	os.Setenv(envLogLevel, "debug")
	defer os.Unsetenv(envLogLevel)
	assert.EqualValues(t, zap.DebugLevel, getLevel())
}

func TestGetLevelInfo(t *testing.T) {
	os.Setenv(envLogLevel, "info")
	defer os.Unsetenv(envLogLevel)
	assert.EqualValues(t, zap.InfoLevel, getLevel())
}

func TestGetLevelError(t *testing.T) {
	os.Setenv(envLogLevel, "error")
	defer os.Unsetenv(envLogLevel)
	assert.EqualValues(t, zap.ErrorLevel, getLevel())
}

func TestGetLevelDefault(t *testing.T) {
	os.Setenv(envLogLevel, "test")
	defer os.Unsetenv(envLogLevel)
	assert.EqualValues(t, zap.InfoLevel, getLevel())
}

func TestGetOutputStdout(t *testing.T) {
	assert.EqualValues(t, "stdout", getOutput())
}

func TestGetOutputSetOutput(t *testing.T) {
	os.Setenv(envLogOutput, "test")
	defer os.Unsetenv(envLogOutput)
	assert.EqualValues(t, "test", getOutput())
}

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	assert.EqualValues(t, "", log.Print("test"))
}
