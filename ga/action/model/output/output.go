package output

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/store"
)

type ActionOutput struct {
	*io.ActionProperty
	TargetReferences *store.PropertyStore[io.PartialActionReference]
}

func NewActionOutput(name string, actionUid string) *ActionOutput {
	return &ActionOutput{
		ActionProperty:   io.NewActionProperty(actionUid, "output", name),
		TargetReferences: store.NewPropertyStore[io.PartialActionReference](true),
	}
}

func (ao *ActionOutput) AssignTargetReference(ref *io.PartialActionReference) {
	ao.TargetReferences.NewProperty(*ref)
}
