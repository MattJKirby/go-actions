package output

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionOutput struct {
	*common.ModelProperty
	TargetReferences *store.ResourceStore[common.ResourceReference] `json:"references"`
}

func NewActionOutput(actionUid uid.ResourceUid, name string) *ActionOutput {
	return &ActionOutput{
		ModelProperty:    common.NewModelProperty(actionUid, "output", name),
		TargetReferences: store.NewResourceStore[common.ResourceReference](true),
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *common.ResourceReference) error {
	return ao.TargetReferences.NewResource(*ref)
}
