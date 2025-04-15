package model

import (
	"fmt"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/action/model/store"
	"go-actions/ga/app/config"
)

type ActionModel struct {
	globalConfig *config.GlobalConfig
	ActionName string                                           `json:"name"`
	ActionUid  string                                           `json:"uid"`
	Parameters *store.PropertyStore[store.IdentifiableProperty] `json:"parameters"`
	Inputs     *store.PropertyStore[input.ActionInput]          `json:"inputs"`
	Outputs    *store.PropertyStore[output.ActionOutput]        `json:"outputs"`
}

func NewActionModel(typename string, globalConfig *config.GlobalConfig) *ActionModel {
	return &ActionModel{
		globalConfig: globalConfig,
		ActionName: typename,
		ActionUid:  fmt.Sprintf("%s:%s", typename, globalConfig.UidGenerator.GenerateUid()),
		Parameters: store.NewPropertyStore[store.IdentifiableProperty](false),
		Inputs:     store.NewPropertyStore[input.ActionInput](false),
		Outputs:    store.NewPropertyStore[output.ActionOutput](false),
	}
}

func Parameter[T any](m *ActionModel, name string, defaultValue T) *parameter.ActionParameter[T] {
	parameterFn := func() *store.IdentifiableProperty {
		value := store.IdentifiableProperty(parameter.NewActionParameter(m.ActionUid, name, defaultValue))
		return &value
	}
	return (*m.Parameters.GetDefault(name, parameterFn)).(*parameter.ActionParameter[T])
}

func Input(m *ActionModel, name string, required bool, defaultSource *output.ActionOutput) *input.ActionInput {
	input := m.Inputs.GetDefault(name, func() *input.ActionInput {
		return input.NewActionInput(name, m.ActionUid)
	})

	if defaultSource != nil {
		io.NewActionReference(m.globalConfig, defaultSource, input).AssignReferences()
	}

	return input
}

func Output(m *ActionModel, name string, defaultTargets []*input.ActionInput) *output.ActionOutput {
	output := m.Outputs.GetDefault(name, func() *output.ActionOutput {
		return output.NewActionOutput(name, m.ActionUid)
	})

	for _, target := range defaultTargets {
		if target != nil {
			io.NewActionReference(m.globalConfig, output, target).AssignReferences()
		}
	}
	return output
}