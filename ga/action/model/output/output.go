package output

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/libs/store"
	"go-actions/ga/libs/uid"
)

type ActionOutput struct {
	*io.ActionProperty
	TargetReferences *store.ResourceStore[io.ActionReference]
}

func NewActionOutput(actionUid *uid.ResourceUid, name string) *ActionOutput {
	return &ActionOutput{
		ActionProperty:   io.NewActionProperty(actionUid, "output", name),
		TargetReferences: store.NewResourceStore[io.ActionReference](true),
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *io.ActionReference) error {
	return ao.TargetReferences.NewResource(*ref)
}
