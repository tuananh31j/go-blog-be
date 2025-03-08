package logger_test

import (
	"testing"

	"nta-blog/components/logger"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	log := logger.NewLogger()

	assert.NotNil(t, log, "Logger should not be nil")
	assert.IsType(t, &zerolog.Logger{}, log, "Logger should be of type *zerolog.Logger")

	// Test logger level
	assert.Equal(t, zerolog.InfoLevel, log.GetLevel(), "Logger level should be Info")

	// Test logger output
	expectedOutput := "test message"
	log.Info().Msg(expectedOutput)
}
