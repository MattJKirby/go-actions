package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/parameter"

	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("test")
	model := model.NewActionModel("test", &model.ModelConfig{})
	model.ActionUid = instance.Model.ActionUid

	asserts.Equals(t, model, instance.Model)
}

func TestParameter(t *testing.T) {
	instance := NewActionInstance("test")
	expected := Parameter(instance, "paramName", 0)

	param, err := instance.Model.Parameters.Get("paramName")
	asserts.Equals(t, nil, err)
	asserts.Equals(t, expected, any(*param).(*parameter.ActionParameter[int]))
}

func TestInput(t *testing.T) {
	instance := NewActionInstance("test")
	expected := Input(instance, "inputName", false)

	input, err := instance.Model.Inputs.Get("inputName")
	asserts.Equals(t, nil, err)
	asserts.Equals(t, expected, input)
}

func TestOutput(t *testing.T) {
	instance := NewActionInstance("test")
	expected := Output(instance, "outputName")

	output, err := instance.Model.Outputs.Get("outputName")
	asserts.Equals(t, nil, err)
	asserts.Equals(t, expected, output)
}
