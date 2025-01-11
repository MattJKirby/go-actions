package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
)

type ActionInstance struct {
	Model *model.ActionModel
}

func NewActionInstance(actionName string) *ActionInstance {
	return &ActionInstance{
		Model: model.NewActionModel(actionName, config.NewModelConfig()),
	}
}

func Parameter[T any](a *ActionInstance, name string, defaultValue T)  *parameter.ActionParameter[T] {
	parameterFn := func() *any {
		value := any(parameter.NewActionParameter(name, defaultValue))
		return &value
	}
	return (*a.Model.Parameters.GetOrDefault(name, parameterFn)).(*parameter.ActionParameter[T])
}

func Input(a *ActionInstance, name string, required bool) *io.Input {
	inputFn := func() *io.Input {
		return io.NewInput(name, a.Model.ActionUid, required)
	}
	return a.Model.Inputs.GetOrDefault(name, inputFn)
}

func Output(a *ActionInstance, name string) *io.Output {
	outputFn := func() *io.Output {
		return io.NewActionOutput(name, a.Model.ActionUid)
	}
	return a.Model.Outputs.GetOrDefault(name, outputFn)
}
