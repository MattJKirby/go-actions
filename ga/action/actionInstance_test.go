package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/action/model/references"
	"go-actions/ga/testing/testHelpers/actionTestHelpers"

	"go-actions/ga/cr/asserts"
	"testing"
)

var mockConfig = &actionTestHelpers.MockActionConfig{MockUid: "uid"}

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("test", mockConfig)
	expectedModel := model.NewActionModel("test", mockConfig)

	asserts.Equals(t, expectedModel, instance.Model)
}

func TestParameter(t *testing.T) {
	instance := NewActionInstance("test", mockConfig)
	expected := Parameter(instance, "paramName", 0)

	param, err := instance.Model.Parameters.Get("paramName")
	asserts.Equals(t, nil, err)
	asserts.Equals(t, expected, any(*param).(*parameter.ActionParameter[int]))
}

func TestInput(t *testing.T) {
	testcases := []struct {
		name                    string
		defaultSource           *output.ActionOutput
		expectedSourceReference *references.ActionReference
	}{
		{name: "without default source", defaultSource: nil},
		{name: "with default source", defaultSource: output.NewActionOutput("o", "a")},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			expected := input.NewActionInput("inputName", instance.Model.ActionUid)

			input := Input(instance, "inputName", false, test.defaultSource)

			asserts.Equals(t, expected, input)
		})
	}
}

func TestOutput(t *testing.T) {
	testcases := []struct {
		name     string
		defaults []*input.ActionInput
		expected []*references.ActionReference
	}{
		{name: "without default targets", defaults: []*input.ActionInput{}, expected: []*references.ActionReference{}},
		{name: "with default targets", defaults: []*input.ActionInput{}, expected: []*references.ActionReference{}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			expected := output.NewActionOutput("outputName", instance.Model.ActionUid)

			// for _, target := range test.expected {
			// 	expected.AssignTarget(target)
			// }

			output := Output(instance, "outputName", test.defaults)

			asserts.Equals(t, expected, output)
		})
	}
}
