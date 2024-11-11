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
		asserts.Equals(t, parameter.value, parameter.Value())
		asserts.Equals(t, parameter.defaultValue, parameter.DefaultValue())
	})

	t.Run("test set parameter", func(t *testing.T) {
		defaultVal := "some string"
		newVal := "test"
		param := NewActionParameter("test", defaultVal)
		param.SetValue(newVal)
		asserts.Equals(t, newVal, param.Value())
		asserts.Equals(t, defaultVal, param.DefaultValue())
	})
}
