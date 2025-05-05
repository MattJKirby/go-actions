package input

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionInput struct {
	*io.ActionProperty
	SourceReferences *store.ResourceStore[io.ActionReference] `json:"references"`
}

func NewActionInput(modelUid *uid.ResourceUid, name string) *ActionInput {
	return &ActionInput{
		ActionProperty:   io.NewActionProperty(modelUid, "input", name),
		SourceReferences: store.NewResourceStore[io.ActionReference](true),
	}
}

func (ai *ActionInput) AssignSourceReference(ref *io.ActionReference) error {
	return ai.SourceReferences.NewResource(*ref)
}
