package output

import (
	"fmt"
	"go-actions/ga/action/model/references"
)

type ActionOutput struct {
	Uid              string
	Name             string
	TargetReferences map[string]*references.ActionReference
}

func NewActionOutput(name string, actionUid string) *ActionOutput {
	return &ActionOutput{
		Uid:              fmt.Sprintf("%s:output:%s", actionUid, name),
		Name:             name,
		TargetReferences: map[string]*references.ActionReference{},
	}
}

func (ao ActionOutput) GetPropertyId() string {
	return ao.Uid
}

func (ao *ActionOutput) AssignTargetReference(ref *references.ActionReference) {
	ao.TargetReferences[ref.ReferenceUid] = ref
}
