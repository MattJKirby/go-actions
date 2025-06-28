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
	ModelUid     uid.ResourceUid                           `json:"uid"`
	Parameters   *store.ResourceStore[store.Identifiable]  `json:"parameters"`
	Inputs       *store.ResourceStore[input.ActionInput]   `json:"inputs"`
	Outputs      *store.ResourceStore[output.ActionOutput] `json:"outputs"`
}

func NewActionModel(globalConfig *config.GlobalConfig, actionUid uid.ResourceUid) *ActionModel {
	return &ActionModel{
		globalConfig: globalConfig,
		ModelUid:     uid.NewUidBuilder().FromParent(actionUid).WithSubResource("Model").Build(),
		Parameters:   store.NewResourceStore(store.Identifiable.GetId, false),
		Inputs:       store.NewResourceStore(input.ActionInput.GetInputId, false),
		Outputs:      store.NewResourceStore(output.ActionOutput.GetOutputId, false),
	}
}

func Parameter[T any](m *ActionModel, name string, defaultValue T) *parameter.ActionParameter[T] {
	defaultFn := func () store.Identifiable {
		return parameter.NewActionParameter(m.ModelUid, name, defaultValue)
	}
	return m.Parameters.GetDefault(name, defaultFn).(*parameter.ActionParameter[T])
}

func Input(m *ActionModel, name string, required bool, source *output.ActionOutput) *input.ActionInput {
	defaultFn := func () input.ActionInput {
		return *input.NewActionInput(m.ModelUid, name)
	}
	input := m.Inputs.GetDefault(name, defaultFn)

	if source != nil {
		Reference(m.globalConfig, source, &input)
	}

	return &input
}

func Output(m *ActionModel, name string, targets []*input.ActionInput) *output.ActionOutput {
	defaultFn := func () output.ActionOutput {
		return *output.NewActionOutput(m.ModelUid, name)
	}

	output := m.Outputs.GetDefault(name, defaultFn)

	for _, target := range targets {
		if target != nil {
			Reference(m.globalConfig, &output, target)
		}
	}
	return &output
}

func Reference(globalConfig *config.GlobalConfig, source *output.ActionOutput, target *input.ActionInput) *common.ResourceReference {
	ref := common.NewActionReference(globalConfig, &source.Uid, &target.Uid)
	source.AssignTargetReference(ref.GetTargetReference())
	target.AssignSourceReference(ref.GetSourceReference())
	return ref
}
