package quickfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logWithTrace(t *testing.T) {
	result := logWithTrace("test_msg")
	expected := "test_msg (log_test.go:10.Test_logWithTrace)"
	assert.Equal(t, expected, result)
}

func Test_logWithTracef(t *testing.T) {
	result := logWithTracef("test_msg_%s", "formatter")
	expected := "test_msg_formatter (log_test.go:16.Test_logWithTracef)"
	assert.Equal(t, expected, result)
}
