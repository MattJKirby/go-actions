package references

import (
	"fmt"
	"go-actions/ga/action/model/store"
)

type ActionInput struct {
	Uid              string
	Name             string
	ActionUid        string
	SourceReferences *store.PropertyStore[ActionReference]
}

func NewActionInput(name string, actionUid string) *ActionInput {
	return &ActionInput{
		Uid:              fmt.Sprintf("%s:input:%s", actionUid, name),
		Name:             name,
		ActionUid:        actionUid,
		SourceReferences: store.NewPropertyStore[ActionReference](),
	}
}

func (ai *ActionInput) AssignSourceReference(ref *ActionReference) {
	ai.SourceReferences.Add(ref.ReferenceUid, ref)
}
