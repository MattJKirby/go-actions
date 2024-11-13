package parameter

import (
	"encoding/json"
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

func TestMarshallParameter(t *testing.T){
	parameter := NewActionParameter("parameterName", "defaultVal")

	t.Run("test marshall parameter", func(t *testing.T) {
		marshalled, err := json.Marshal(parameter)
		if err != nil {
			t.Errorf("error marshalling parameter: got %v", err)
		}

		asserts.Equals(t, `{"Name":"parameterName","Value":"defaultVal"}`, string(marshalled))
	})
}
