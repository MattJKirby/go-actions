package output

import (
	"go-actions/ga/action/model/io"
)

type ActionOutput struct {
	*io.ActionProperty
	TargetReferences map[string]*io.ActionReference
}

func NewActionOutput(name string, actionUid string) *ActionOutput {
	return &ActionOutput{
		ActionProperty:   io.NewActionProperty(actionUid, "output", name),
		TargetReferences: map[string]*io.ActionReference{},
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *io.ActionReference) {
	ao.TargetReferences[ref.ReferenceUid] = ref
}
