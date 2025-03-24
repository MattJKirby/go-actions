package references

import "fmt"

type ActionInput struct {
	Uid              string
	Name             string
	ActionUid        string
	SourceReferences map[string]*ActionReference
}

func NewActionInput(name string, actionUid string) *ActionInput {
	return &ActionInput{
		Uid:              fmt.Sprintf("%s:input:%s", actionUid, name),
		Name:             name,
		ActionUid:        actionUid,
		SourceReferences: map[string]*ActionReference{},
	}
}

func (ai *ActionInput) AssignSourceReference(ref *ActionReference) {
	ai.SourceReferences[ref.ReferenceUid] = ref
}
