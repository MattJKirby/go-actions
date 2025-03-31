package input

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/store"
)

type ActionInput struct {
	*io.ActionProperty
	SourceReferences *store.PropertyStore[io.ActionReference]
}

func NewActionInput(name string, actionUid string) *ActionInput {
	return &ActionInput{
		ActionProperty: io.NewActionProperty(actionUid,"input",name),
		SourceReferences: store.NewPropertyStore[io.ActionReference](true),
	}
}

func (ai *ActionInput) AssignSourceReference(ref *io.ActionReference) {
	ai.SourceReferences.NewProperty(*ref)
}
