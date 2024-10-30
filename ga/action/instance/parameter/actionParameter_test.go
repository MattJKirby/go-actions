package parameter

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewParameter(t *testing.T) {
	parameter := NewActionParameter("test", "default value")

	t.Run("test new parameter", func(t *testing.T) {
		asserts.Equals(t, "test", parameter.name)
		asserts.Equals(t, "default value", parameter.value)
		asserts.Equals(t, "default value", parameter.defaultValue)
	})
}