package input

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionInput struct {
	*common.ModelProperty
	SourceReferences *store.ResourceStore[common.ActionReference] `json:"references"`
}

func NewActionInput(modelUid uid.ResourceUid, name string) *ActionInput {
	return &ActionInput{
		ModelProperty:    common.NewModelProperty(modelUid, "input", name),
		SourceReferences: store.NewResourceStore[common.ActionReference](true),
	}
}

func (ai *ActionInput) AssignSourceReference(ref *common.ActionReference) error {
	return ai.SourceReferences.NewResource(*ref)
}
