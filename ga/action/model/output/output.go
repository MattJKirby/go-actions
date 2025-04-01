package output

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/store"
)

type ActionOutput struct {
	*io.ActionProperty
	TargetReferences *store.PropertyStore[io.ActionReference]
}

func NewActionOutput(name string, actionUid string) *ActionOutput {
	return &ActionOutput{
		ActionProperty:   io.NewActionProperty(actionUid, "output", name),
		TargetReferences: store.NewPropertyStore[io.ActionReference](true),
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *io.ActionReference) {
  ao.TargetReferences.NewProperty(*ref)
}
