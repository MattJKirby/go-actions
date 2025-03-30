package references

import "fmt"

type ActionOutput struct {
	Uid              string
	Name             string
	actionUid        string
	TargetReferences map[string]*ActionReference
}

func NewActionOutput(name string, actionUid string) *ActionOutput {
	return &ActionOutput{
		Uid:              fmt.Sprintf("%s:output:%s", actionUid, name),
		Name:             name,
		actionUid:        actionUid,
		TargetReferences: map[string]*ActionReference{},
	}
}

func (ao ActionOutput) GetPropertyId() string {
	return ao.Uid
}

func (ao *ActionOutput) AssignTargetReference(ref *ActionReference) {
	ao.TargetReferences[ref.ReferenceUid] = ref
}
