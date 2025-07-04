package output

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionOutput struct {
	Name             string                                         `json:"name"`
	Uid              uid.ResourceUid                                `json:"uid"`
	TargetReferences *store.ResourceStore[common.ResourceReference] `json:"references"`
}

func NewActionOutput(actionUid uid.ResourceUid, name string) *ActionOutput {
	return &ActionOutput{
		Name:             name,
		Uid:              uid.NewUidBuilder().FromParent(actionUid).WithSubResource("output").WithSubResourceId(name).Build(),
		TargetReferences: store.NewResourceStore(common.ResourceReference.GetId, true),
	}
}

func (ao ActionOutput) GetOutputId() string {
	return ao.Uid.FullUid()
}

func (ao *ActionOutput) AssignTargetReference(ref *common.ResourceReference) error {
	return ao.TargetReferences.NewResource(*ref)
}
