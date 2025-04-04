package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"

	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("test", mockConfig)
	expectedModel := model.NewActionModel("test", mockConfig)

	assert.Equals(t, expectedModel, instance.Model)
}

func TestParameter(t *testing.T) {
	instance := NewActionInstance("test", mockConfig)
	expected := Parameter(instance, "paramName", 0)

	param, err := instance.Model.Parameters.Get("paramName")
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, any(*param).(*parameter.ActionParameter[int]))
}

func TestInput(t *testing.T) {
	testcases := []struct {
		name           string
		defaultSource  *output.ActionOutput
		expectedSource *io.ActionReference
	}{
		{name: "without default source", defaultSource: nil, expectedSource: nil},
		{name: "with default source", defaultSource: output.NewActionOutput("o", "a")},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			expected := input.NewActionInput("inputName", instance.Model.ActionUid)

			input := Input(instance, "inputName", false, test.defaultSource)

			assert.Equals(t, expected, input)
		})
	}
}

func TestOutput(t *testing.T) {
	testcases := []struct {
		name     string
		defaults []*input.ActionInput
		expected []*io.ActionReference
	}{
		{name: "without default targets", defaults: []*input.ActionInput{}, expected: []*io.ActionReference{}},
		{name: "with default targets", defaults: []*input.ActionInput{}, expected: []*io.ActionReference{}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			expected := output.NewActionOutput("outputName", instance.Model.ActionUid)

			// for _, target := range test.expected {
			// 	expected.AssignTarget(target)
			// }

			output := Output(instance, "outputName", test.defaults)

			assert.Equals(t, expected, output)
		})
	}
}
