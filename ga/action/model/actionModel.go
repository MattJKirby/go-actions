package model

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/action/model/input"
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
	value := parameter.NewActionParameter(m.ModelUid, name, defaultValue)
	return m.Parameters.GetDefault(value).(*parameter.ActionParameter[T])
}

func Input(m *ActionModel, name string, required bool, source *output.ActionOutput) *input.ActionInput {
	value := input.NewActionInput(m.ModelUid, name)
	input := m.Inputs.GetDefault(*value)

	if source != nil {
		Reference(m.globalConfig, source, &input)
	}

	return &input
}

func Output(m *ActionModel, name string, targets []*input.ActionInput) *output.ActionOutput {
	value := output.NewActionOutput(m.ModelUid, name)
	output := m.Outputs.GetDefault(*value)

	for _, target := range targets {
		if target != nil {
			Reference(m.globalConfig, &output, target)
		}
	}
	return &output
}

func Reference(globalConfig *config.GlobalConfig, source *output.ActionOutput, target *input.ActionInput) *common.ActionReference {
	ref := common.NewActionReference(globalConfig, source.Uid, target.Uid)
	source.AssignTargetReference(ref)
	target.AssignSourceReference(ref)
	return ref
}
