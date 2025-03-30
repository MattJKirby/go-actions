package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/action/model/references"
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
	parameterFn := func() *any {
		value := any(parameter.NewActionParameter(name, defaultValue))
		return &value
	}
	return (*a.Model.Parameters.GetOrDefault(name, parameterFn)).(*parameter.ActionParameter[T])
}

func Input(a *ActionInstance, name string, required bool, defaultSource *references.ActionOutput) *references.ActionInput {
	inputFn := func() *references.ActionInput {
		return references.NewActionInput(name, a.Model.ActionUid)
	}

	input := a.Model.Inputs.GetDefault(name, inputFn)
	return input
}

func Output(a *ActionInstance, name string, defaultTargets []*references.ActionInput) *references.ActionOutput {
	outputFn := func() *references.ActionOutput {
		return references.NewActionOutput(name, a.Model.ActionUid)
	}
	output := a.Model.Outputs.GetDefault(name, outputFn)
	return output
}
