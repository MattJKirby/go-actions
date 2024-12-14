package model

import (
	"fmt"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
)

type ActionModel struct {
	ActionName string               `json:"name"`
	ActionUid  string               `json:"uid"`
	Parameters *parameter.Store     `json:"parameters"`
	Inputs     *io.Store[io.Input]  `json:"inputs"`
	Outputs    *io.Store[io.Output] `json:"outputs"`
}

type ActionModelConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionModelConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: parameter.NewStore(),
		Inputs:     io.NewStore[io.Input](),
		Outputs:    io.NewStore[io.Output](),
	}
}

func Parameter[T any](name string, defaultValue T) func(*ActionModel) *parameter.ActionParameter[T] {
	return func(m *ActionModel) *parameter.ActionParameter[T] {
		defaultAsAny := any(parameter.NewActionParameter(name, defaultValue))
		got := m.Parameters.GetOrDefault(name, &defaultAsAny)
		return (*got).(*parameter.ActionParameter[T])
	}
}

func Input(name string, required bool) func(*ActionModel) *io.Input {
	return func(m *ActionModel) *io.Input {
		defaultInput := io.NewInput(name, m.ActionUid, required)
		return m.Inputs.GetOrDefault(name, defaultInput)
	}
}

func Output(name string) func(*ActionModel) *io.Output {
	return func(m *ActionModel) *io.Output {
		defaultOutput := io.NewActionOutput(name, m.ActionUid)
		return m.Outputs.GetOrDefault(name, defaultOutput)
	}
}
