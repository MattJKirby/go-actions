package output

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionOutput struct {
	*common.ModelProperty
	TargetReferences *store.ResourceStore[common.ActionReference] `json:"references"`
}

func NewActionOutput(actionUid uid.ResourceUid, name string) *ActionOutput {
	return &ActionOutput{
		ModelProperty:   common.NewModelProperty(actionUid, "output", name),
		TargetReferences: store.NewResourceStore[common.ActionReference](true),
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *common.ActionReference) error {
	return ao.TargetReferences.NewResource(*ref)
}
