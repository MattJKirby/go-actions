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
		name          string
		defaultSource *output.ActionOutput
		ref           *io.PartialActionReference
	}{
		{name: "without default source"},
		{name: "with default source", defaultSource: output.NewActionOutput("o", "a"), ref: &io.PartialActionReference{ReferenceUid: "ref:uid", ActionUid: "a"}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			input := Input(instance, "inputName", false, test.defaultSource)

			expectedInput, err := instance.Model.Inputs.Get("inputName")
			ref, _ := input.SourceReferences.Get("ref:uid")

			assert.Equals(t, expectedInput, input)
			assert.Equals(t, nil, err)
			assert.Equals(t, test.ref, ref)
		})
	}
}

func TestOutput(t *testing.T) {
	// testcases := []struct {
	// 	name     string
	// 	defaults []*input.ActionInput
	// 	expected []*io.ActionReference
	// }{
	// 	{name: "without default targets", defaults: []*input.ActionInput{}, expected: []*io.ActionReference{}},
	// 	{name: "with default targets", defaults: []*input.ActionInput{}, expected: []*io.ActionReference{}},
	// }

	// for _, test := range testcases {
	// 	t.Run(test.name, func(t *testing.T) {
	// 		instance := NewActionInstance("test", mockConfig)
	// 		expected := output.NewActionOutput("outputName", instance.Model.ActionUid)

	// 		// for _, target := range test.expected {
	// 		// 	expected.AssignTarget(target)
	// 		// }

	// 		output := Output(instance, "outputName", test.defaults)

	// 		assert.Equals(t, expected, output)
	// 	})
	// }

	testcases := []struct {
		name          string
		defaultTargets []*input.ActionInput
		ref           *io.PartialActionReference
	}{
		{name: "without default source", defaultTargets: make([]*input.ActionInput, 0)},
		{name: "with default source", defaultTargets: []*input.ActionInput{input.NewActionInput("1", "a")}, ref: &io.PartialActionReference{ReferenceUid: "ref:uid", ActionUid: "a"}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			instance := NewActionInstance("test", mockConfig)
			output := Output(instance, "outputName", test.defaultTargets)

			expectedOutput, err := instance.Model.Outputs.Get("outputName")
			ref, _ := output.TargetReferences.Get("ref:uid")

			assert.Equals(t, expectedOutput, output)
			assert.Equals(t, nil, err)
			assert.Equals(t, test.ref, ref)
		})
	}
}
