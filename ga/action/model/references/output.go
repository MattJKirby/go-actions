package references

import "fmt"

type ActionOutput struct {
	Uid              string
	Name             string
	ActionUid        string
	TargetReferences map[string]*ActionReference
}

func NewActionOutput(name string, actionUid string) *ActionOutput {
	return &ActionOutput{
		Uid:              fmt.Sprintf("%s:input:%s", actionUid, name),
		Name:             name,
		ActionUid:        actionUid,
		TargetReferences: map[string]*ActionReference{},
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *ActionReference){
	ao.TargetReferences[ref.ReferenceUid] = ref
}