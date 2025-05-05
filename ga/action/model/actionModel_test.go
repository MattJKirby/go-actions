package model

import (
	"encoding/json"
	"go-actions/ga/action/model/input"
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
	model := NewActionModel(mockConfig, uid)
	mashalled, _ := json.Marshal(model)

	assert.Equals(t, `{"uid":"::::model:","parameters":[],"inputs":[],"outputs":[]}`, string(mashalled))
}

func TestUnmarshalModel(t *testing.T) {
	uid := &uid.ResourceUid{}
	model := NewActionModel(mockConfig, uid)
	marshalled := `{"uid":"::::x:","parameters":[],"inputs":[],"outputs":[]}`

	err := json.Unmarshal([]byte(marshalled), model)
	assert.Equals(t, err, nil)
	assert.Equals(t, model.ModelUid.GetUid(), "::::x:")
}

func TestParameter(t *testing.T) {
	uid := &uid.ResourceUid{}
	model := NewActionModel(mockConfig, uid)
	expected := Parameter(model, "paramName", 0)

	param, err := model.Parameters.Get("paramName")
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, any(*param).(*parameter.ActionParameter[int]))
}

func TestInput(t *testing.T) {
	uid := &uid.ResourceUid{}
	
	testcases := []struct {
		name          string
		defaultSource *output.ActionOutput
	}{
		{name: "without default source"},
		{name: "with default source", defaultSource: output.NewActionOutput(uid, "a")},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			
			model := NewActionModel(mockConfig, uid)
			input := Input(model, "inputName", false, test.defaultSource)

			expectedInput, err := model.Inputs.Get("inputName")

			assert.Equals(t, expectedInput, input)
			assert.Equals(t, nil, err)
		})
	}
}

func TestOutput(t *testing.T) {
	uid := &uid.ResourceUid{}

	testcases := []struct {
		name           string
		defaultTargets []*input.ActionInput
	}{
		{name: "without default source", defaultTargets: make([]*input.ActionInput, 0)},
		{name: "with default source", defaultTargets: []*input.ActionInput{input.NewActionInput(uid, "a")}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			model := NewActionModel(mockConfig, uid)
			output := Output(model, "outputName", test.defaultTargets)

			expectedOutput, err := model.Outputs.Get("outputName")

			assert.Equals(t, expectedOutput, output)
			assert.Equals(t, nil, err)
		})
	}
}
