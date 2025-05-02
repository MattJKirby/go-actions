package model

import (
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/app/config"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionModel struct {
	globalConfig *config.GlobalConfig
	ModelUid    *uid.ResourceUid                                       `json:"uid"`
	Parameters   *store.ResourceStore[store.IdentifiableResource] `json:"parameters"`
	Inputs       *store.ResourceStore[input.ActionInput]          `json:"inputs"`
	Outputs      *store.ResourceStore[output.ActionOutput]        `json:"outputs"`
}

func NewActionModel(actionUid *uid.ResourceUid, globalConfig *config.GlobalConfig) *ActionModel {
	return &ActionModel{
		globalConfig: globalConfig,
		ModelUid:    actionUid.FromParent(uid.WithSubResource("model")),
		Parameters:   store.NewResourceStore[store.IdentifiableResource](false),
		Inputs:       store.NewResourceStore[input.ActionInput](false),
		Outputs:      store.NewResourceStore[output.ActionOutput](false),
	}
}

func Parameter[T any](m *ActionModel, name string, defaultValue T) *parameter.ActionParameter[T] {
	parameterFn := func() *store.IdentifiableResource {
		value := store.IdentifiableResource(parameter.NewActionParameter(m.ModelUid.GetString(), name, defaultValue))
		return &value
	}
	return (*m.Parameters.GetDefault(name, parameterFn)).(*parameter.ActionParameter[T])
}

func Input(m *ActionModel, name string, required bool, defaultSource *output.ActionOutput) *input.ActionInput {
	input := m.Inputs.GetDefault(name, func() *input.ActionInput {
		return input.NewActionInput(name, m.ModelUid.GetString())
	})

	if defaultSource != nil {
		io.NewActionReference(m.globalConfig, defaultSource, input).AssignReferences()
	}

	return input
}

func Output(m *ActionModel, name string, defaultTargets []*input.ActionInput) *output.ActionOutput {
	output := m.Outputs.GetDefault(name, func() *output.ActionOutput {
		return output.NewActionOutput(name, m.ModelUid.GetString())
	})

	for _, target := range defaultTargets {
		if target != nil {
			io.NewActionReference(m.globalConfig, output, target).AssignReferences()
		}
	}
	return output
}
