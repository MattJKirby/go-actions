package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
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
		defaultSource           *io.Output
		expectedSourceReference *io.ActionReference
	}{
		{name: "without default source", defaultSource: nil},
		{name: "with default source", defaultSource: io.NewActionOutput("o", "a"), expectedSourceReference: io.NewReference("a", "o")},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			expected := io.NewInput("inputName", instance.Model.ActionUid, false)
			expected.AssignSource(test.expectedSourceReference)

			input := Input(instance, "inputName", false, test.defaultSource)

			asserts.Equals(t, expected, input)
		})
	}
}

func TestOutput(t *testing.T) {
	testcases := []struct {
		name     string
		defaults []*io.Input
		expected []*io.ActionReference
	}{
		{name: "without default targets", defaults: []*io.Input{}, expected: []*io.ActionReference{}},
		{name: "with default targets", defaults: []*io.Input{io.NewInput("i", "a", false), io.NewInput("i", "b", false)}, expected: []*io.ActionReference{io.NewReference("a", "i"), io.NewReference("b", "i")}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			expected := io.NewActionOutput("outputName", instance.Model.ActionUid)

			for _, target := range test.expected {
				expected.AssignTarget(target)
			}

			output := Output(instance, "outputName", test.defaults)

			asserts.Equals(t, expected, output)
		})
	}
}
