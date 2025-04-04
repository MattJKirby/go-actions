package parameter

import (
	"encoding/json"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestNewParameter(t *testing.T) {
	parameter := NewActionParameter("uid", "test", "default value")

	t.Run("test new parameter", func(t *testing.T) {
		assert.Equals(t, "test", parameter.Name)
		assert.Equals(t, "default value", parameter.value)
		assert.Equals(t, "default value", parameter.defaultValue)
		assert.Equals(t, parameter.value, parameter.Value())
		assert.Equals(t, parameter.defaultValue, parameter.DefaultValue())
	})

	t.Run("test set parameter", func(t *testing.T) {
		defaultVal := "some string"
		newVal := "test"
		param := NewActionParameter("uid", "test", defaultVal)
		param.SetValue(newVal)
		assert.Equals(t, newVal, param.Value())
		assert.Equals(t, defaultVal, param.DefaultValue())
	})
}

func TestMarshalParameter(t *testing.T) {
	parameter := NewActionParameter("uid", "parameterName", "defaultVal")
	expectedMarshalResult := `{"Uid":"uid:parameter:parameterName","name":"parameterName","value":"defaultVal"}`

	marshalled, err := json.Marshal(parameter)
	if err != nil {
		t.Errorf("error marshalling parameter: got %v", err)
	}

	assert.Equals(t, expectedMarshalResult, string(marshalled))
}

func TestUnmarshalParameter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      bool
	}{
		{name: "unmarshal valid marshalled input", input: `{"Name":"parameterName","Value":"changedVal"}`, expected: "changedVal", err: false},
		{name: "unmarshal invalid marshalled input (bad name)", input: `{"Name":"badName","Value":"changedVal"}`, expected: "defaultVal", err: true},
		{name: "unmarshal invalid marshalled input (bad input)", input: `{"Name":"parameterName","Value":0}`, expected: "defaultVal", err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parameter := NewActionParameter("uid", "parameterName", "defaultVal")
			err := json.Unmarshal([]byte(test.input), parameter)

			if test.err != (err != nil) {
				t.Errorf("error unmarshalling parameter: got %v", err)
			}

			assert.Equals(t, test.expected, parameter.value)
		})
	}
}
