package references

import (
	"fmt"
	"go-actions/ga/action/model/store"
)

type ActionInput struct {
	Uid              string
	Name             string
	actionUid        string
	SourceReferences *store.ActionPropertyStore[ActionReference]
}

func NewActionInput(name string, actionUid string) *ActionInput {
	return &ActionInput{
		Uid:              fmt.Sprintf("%s:input:%s", actionUid, name),
		Name:             name,
		actionUid:        actionUid,
		SourceReferences: store.NewActionPropertyStore[ActionReference](true),
	}
}

func (ai ActionInput) GetPropertyId() string {
	return ai.Uid
}

func (ai *ActionInput) AssignSourceReference(ref *ActionReference) {
	ai.SourceReferences.NewProperty(*ref)
}
