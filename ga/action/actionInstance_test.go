package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/io"
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
			instance := NewActionInstance("test")
			expected := io.NewInput("inputName", instance.Model.ActionUid, false)
			expected.AssignSource(test.expectedSourceReference)

			input := Input(instance, "inputName", false, test.defaultSource)

			asserts.Equals(t, expected, input)
		})
	}

}

func TestOutput(t *testing.T) {
	instance := NewActionInstance("test")
	expected := Output(instance, "outputName")

	output, err := instance.Model.Outputs.Get("outputName")
	asserts.Equals(t, nil, err)
	asserts.Equals(t, expected, output)
}
