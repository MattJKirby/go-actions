package parameter

import (
	"encoding/json"
	"go-actions/ga/cr"
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

func TestMarshalParameter(t *testing.T) {
	parameter := NewActionParameter("parameterName", "defaultVal")
	expectedMarshalResult := `{"name":"parameterName","value":"defaultVal"}`

	t.Run("test custom parameter marshalling", func(t *testing.T) {
		marshalled, err := json.Marshal(parameter)
		if err != nil {
			t.Errorf("error marshalling parameter: got %v", err)
		}

		asserts.Equals(t, expectedMarshalResult, string(marshalled))
	})
}

func TestUnmarshalParameter(t *testing.T) {

	tests := []cr.TestCase[string, string]{
		{Name: "unmarshal valid marshalled input", Input: `{"Name":"parameterName","Value":"changedVal"}`, Expected: "changedVal", Error: false},
		{Name: "unmarshal invalid marshalled input (bad name)", Input: `{"Name":"badName","Value":"changedVal"}`, Expected: "defaultVal", Error: true},
		{Name: "unmarshal invalid marshalled input (bad input)", Input: `{"Name":"parameterName","Value":0}`, Expected: "defaultVal", Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, string]) {
		parameter := NewActionParameter("parameterName", "defaultVal")
		err := json.Unmarshal([]byte(test.Input), parameter)

		if test.Error != (err != nil) {
			t.Errorf("error unmarshalling parameter: got %v", err)
		}

		asserts.Equals(t, test.Expected, parameter.value)
	})
}
