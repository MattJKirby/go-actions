package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/action/model/store"
)

type ActionInstance struct {
	Model *model.ActionModel `json:"model"`
}

func NewActionInstance(actionName string, modelConfig model.ActionConfig) *ActionInstance {
	return &ActionInstance{
		Model: model.NewActionModel(actionName, modelConfig),
	}
}

func Parameter[T any](a *ActionInstance, name string, defaultValue T) *parameter.ActionParameter[T] {
	parameterFn := func() *store.IdentifiableProperty {
		value := store.IdentifiableProperty(parameter.NewActionParameter(a.Model.ActionUid, name, defaultValue))
		return &value
	}
	return (*a.Model.Parameters.GetDefault(name, parameterFn)).(*parameter.ActionParameter[T])
}

func Input(a *ActionInstance, name string, required bool, defaultSource *output.ActionOutput) *input.ActionInput {
	inputFn := func() *input.ActionInput {
		return input.NewActionInput(name, a.Model.ActionUid)
	}

	input := a.Model.Inputs.GetDefault(name, inputFn)
	return input
}

func Output(a *ActionInstance, name string, defaultTargets []*input.ActionInput) *output.ActionOutput {
	outputFn := func() *output.ActionOutput {
		return output.NewActionOutput(name, a.Model.ActionUid)
	}
	output := a.Model.Outputs.GetDefault(name, outputFn)
	return output
}
