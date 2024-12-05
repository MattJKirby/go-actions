package io

import (
	"fmt"
	"go-actions/ga/action/model/io/reference"
)

type ActionOutput struct {
	Name            string                      `json:"name"`
	Id              string                      `json:"id"`
	InputReferences []*reference.InputReference `json:"inputs"`
}

func NewActionOutput(name string, actionUid string) *ActionOutput {
	id := fmt.Sprintf("%s__Output:%s", actionUid, name)
	inputReferences := []*reference.InputReference{}
	return &ActionOutput{
		name,
		id,
		inputReferences,
	}
}

func (ao *ActionOutput) AssignInputReference(ref *reference.InputReference) {
	ao.InputReferences = append(ao.InputReferences, ref)
}
