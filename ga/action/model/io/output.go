package io

import (
	"fmt"
	"go-actions/ga/action/model/io/reference"
	"go-actions/ga/utils/marshalling"
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

func (ao *ActionOutput) UnmarshalJSON(data []byte) (error) {
	type alias ActionOutput
	var temp alias

	if _, err := marshalling.StrictDecode(data, &temp); err != nil {
		return err
	}

	if temp.Name != ao.Name {
		return fmt.Errorf("name mismatch: got '%s', expected '%s'", temp.Name, ao.Name)
	}

	*ao = ActionOutput(temp)
	return nil
}
