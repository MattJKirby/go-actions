package input

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionInput struct {
	Name string `json:"name"`
	Uid uid.ResourceUid `json:"uid"`
	SourceReferences *store.ResourceStore[common.ResourceReference] `json:"references"`
}

func NewActionInput(modelUid uid.ResourceUid, name string) *ActionInput {
	return &ActionInput{
		Name: name,
		Uid:  uid.NewUidBuilder().FromParent(modelUid).WithSubResource("input").WithSubResourceId(name).Build(),
		SourceReferences: store.NewResourceStore(common.ResourceReference.GetId, true),
	}
}

func (ai ActionInput) GetInputId() string {
	return ai.Uid.FullUid()
}

func (ai *ActionInput) AssignSourceReference(ref *common.ResourceReference) error {
	return ai.SourceReferences.NewResource(*ref)
}
