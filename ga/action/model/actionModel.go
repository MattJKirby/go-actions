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
	ModelUid     *uid.ResourceUid                                 `json:"uid"`
	Parameters   *store.ResourceStore[store.IdentifiableResource] `json:"parameters"`
	Inputs       *store.ResourceStore[input.ActionInput]          `json:"inputs"`
	Outputs      *store.ResourceStore[output.ActionOutput]        `json:"outputs"`
}

func NewActionModel(globalConfig *config.GlobalConfig, actionUid *uid.ResourceUid) *ActionModel {
	return &ActionModel{
		globalConfig: globalConfig,
		ModelUid:     uid.NewResourceUid(uid.WithParentUid(actionUid), uid.WithSubResource("Model")),
		Parameters:   store.NewResourceStore[store.IdentifiableResource](false),
		Inputs:       store.NewResourceStore[input.ActionInput](false),
		Outputs:      store.NewResourceStore[output.ActionOutput](false),
	}
}

func Parameter[T any](m *ActionModel, name string, defaultValue T) *parameter.ActionParameter[T] {
	parameterFn := func() *store.IdentifiableResource {
		value := store.IdentifiableResource(parameter.NewActionParameter(m.ModelUid, name, defaultValue))
		return &value
	}
	return (*m.Parameters.GetDefault(name, parameterFn)).(*parameter.ActionParameter[T])
}

func Input(m *ActionModel, name string, required bool, source *output.ActionOutput) *input.ActionInput {
	input := m.Inputs.GetDefault(name, func() *input.ActionInput {
		return input.NewActionInput(m.ModelUid, name)
	})

	if source != nil {
		ref := io.NewActionReference(m.globalConfig, source.Uid, input.Uid)
		source.AssignTargetReference(ref)
		input.AssignSourceReference(ref)
	}

	return input
}

func Output(m *ActionModel, name string, targets []*input.ActionInput) *output.ActionOutput {
	output := m.Outputs.GetDefault(name, func() *output.ActionOutput {
		return output.NewActionOutput(m.ModelUid, name)
	})

	for _, target := range targets {
		if target != nil {
			ref := io.NewActionReference(m.globalConfig, output.Uid, target.Uid)
			output.AssignTargetReference(ref)
			target.AssignSourceReference(ref)
		}
	}
	return output
}
