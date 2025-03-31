package input

import (
	"fmt"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/store"
)

type ActionInput struct {
	Uid              string
	Name             string
	SourceReferences *store.PropertyStore[io.ActionReference]
}

func NewActionInput(name string, actionUid string) *ActionInput {
	return &ActionInput{
		Uid:              fmt.Sprintf("%s:input:%s", actionUid, name),
		Name:             name,
		SourceReferences: store.NewPropertyStore[io.ActionReference](true),
	}
}

func (ai ActionInput) GetPropertyId() string {
	return ai.Uid
}

func (ai *ActionInput) AssignSourceReference(ref *io.ActionReference) {
	ai.SourceReferences.NewProperty(*ref)
}
