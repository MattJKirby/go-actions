package output

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionOutput struct {
	*common.ActionProperty
	TargetReferences *store.ResourceStore[common.ActionReference] `json:"references"`
}

func NewActionOutput(actionUid *uid.ResourceUid, name string) *ActionOutput {
	return &ActionOutput{
		ActionProperty:   common.NewActionProperty(actionUid, "output", name),
		TargetReferences: store.NewResourceStore[common.ActionReference](true),
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *common.ActionReference) error {
	return ao.TargetReferences.NewResource(*ref)
}
