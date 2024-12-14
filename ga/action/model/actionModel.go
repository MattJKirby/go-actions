package model

import (
	"fmt"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/action/model/resources"
)

type ActionModel struct {
	ActionName string                              `json:"name"`
	ActionUid  string                              `json:"uid"`
	Parameters *resources.ResourceStore[any]       `json:"parameters"`
	Inputs     *resources.ResourceStore[io.Input]  `json:"inputs"`
	Outputs    *resources.ResourceStore[io.Output] `json:"outputs"`
}

type ActionModelConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionModelConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: resources.NewResourceStore[any](),
		Inputs:     resources.NewResourceStore[io.Input](),
		Outputs:    resources.NewResourceStore[io.Output](),
	}
}

func Parameter[T any](name string, defaultValue T) func(*ActionModel) *parameter.ActionParameter[T] {
	return func(m *ActionModel) *parameter.ActionParameter[T] {
		parameterFn := func() *any {
			value := any(parameter.NewActionParameter(name, defaultValue))
			return &value
		}
		return (*m.Parameters.GetOrDefault(name, parameterFn)).(*parameter.ActionParameter[T])
	}
}

func Input(name string, required bool) func(*ActionModel) *io.Input {
	return func(m *ActionModel) *io.Input {
		inputFn := func() *io.Input {
			return io.NewInput(name, m.ActionUid, required)
		}
		return m.Inputs.GetOrDefault(name, inputFn)
	}
}

func Output(name string) func(*ActionModel) *io.Output {
	return func(m *ActionModel) *io.Output {
		outputFn := func() *io.Output {
			return io.NewActionOutput(name, m.ActionUid)
		}
		return m.Outputs.GetOrDefault(name, outputFn)
	}
}
