package model

import (
	"encoding/json"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestMarshalEmptyModel(t *testing.T) {
	uid := &uid.ResourceUid{}
	model := NewActionModel(uid, mockConfig)
	mashalled, _ := json.Marshal(model)

	assert.Equals(t, `{"uid":"::::model:","parameters":[],"inputs":[],"outputs":[]}`, string(mashalled))
}

func TestUnmarshalModel(t *testing.T) {
	uid := &uid.ResourceUid{}
	model := NewActionModel(uid, mockConfig)
	marshalled := `{"uid":"::::x:","parameters":[],"inputs":[],"outputs":[]}`

	err := json.Unmarshal([]byte(marshalled), model)
	assert.Equals(t, err, nil)
	assert.Equals(t, model.ModelUid.GetString(), "::::x:")
}

func TestParameter(t *testing.T) {
	uid := &uid.ResourceUid{}
	model := NewActionModel(uid, mockConfig)
	expected := Parameter(model, "paramName", 0)

	param, err := model.Parameters.Get("paramName")
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
			uid := &uid.ResourceUid{}
			model := NewActionModel(uid, mockConfig)
			input := Input(model, "inputName", false, test.defaultSource)

			expectedInput, err := model.Inputs.Get("inputName")
			ref, _ := input.SourceReferences.Get("ref:uid")

			assert.Equals(t, expectedInput, input)
			assert.Equals(t, nil, err)
			assert.Equals(t, test.ref, ref)
		})
	}
}

func TestOutput(t *testing.T) {
	testcases := []struct {
		name           string
		defaultTargets []*input.ActionInput
		ref            *io.PartialActionReference
	}{
		{name: "without default source", defaultTargets: make([]*input.ActionInput, 0)},
		{name: "with default source", defaultTargets: []*input.ActionInput{input.NewActionInput("1", "a")}, ref: &io.PartialActionReference{ReferenceUid: "ref:uid", ActionUid: "a"}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			uid := &uid.ResourceUid{}
			model := NewActionModel(uid, mockConfig)
			output := Output(model, "outputName", test.defaultTargets)

			expectedOutput, err := model.Outputs.Get("outputName")
			ref, _ := output.TargetReferences.Get("ref:uid")

			assert.Equals(t, expectedOutput, output)
			assert.Equals(t, nil, err)
			assert.Equals(t, test.ref, ref)
		})
	}
}
