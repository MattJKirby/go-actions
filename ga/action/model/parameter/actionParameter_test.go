package parameter

import (
	"encoding/json"
	"go-actions/ga/cr"
	"go-actions/ga/testing/assert"
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
	tests := []cr.TestCase[string, string]{
		{Name: "unmarshal valid marshalled input", Input: `{"Name":"parameterName","Value":"changedVal"}`, Expected: "changedVal", Error: false},
		{Name: "unmarshal invalid marshalled input (bad name)", Input: `{"Name":"badName","Value":"changedVal"}`, Expected: "defaultVal", Error: true},
		{Name: "unmarshal invalid marshalled input (bad input)", Input: `{"Name":"parameterName","Value":0}`, Expected: "defaultVal", Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, string]) {
		parameter := NewActionParameter("uid", "parameterName", "defaultVal")
		err := json.Unmarshal([]byte(test.Input), parameter)

		if test.Error != (err != nil) {
			t.Errorf("error unmarshalling parameter: got %v", err)
		}

		assert.Equals(t, test.Expected, parameter.value)
	})
}
