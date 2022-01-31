package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVar_key_exists(t *testing.T) {
	os.Clearenv()
	value := "TEST_VAR_VALUE"
	value_fallback := "TEST_VAR_VALUE_FALLBACK"
	key := "TEST_VAR_KEY"

	os.Setenv(key, value)

	get_value := GetEnvVar(key, value_fallback)

	assert.NotNil(t, get_value)
	assert.Equal(t, get_value, value)
}

func TestGetEnvVar_key_not_exists(t *testing.T) {
	os.Clearenv()
	value_fallback := "TEST_VAR_VALUE_FALLBACK"
	key := "TEST_VAR_KEY"

	get_value := GetEnvVar(key, value_fallback)

	assert.NotNil(t, get_value)
	assert.Equal(t, get_value, value_fallback)
}
