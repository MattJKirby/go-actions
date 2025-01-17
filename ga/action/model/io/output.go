package io

import (
	"fmt"
	"go-actions/ga/utils/marshalling"
)

type Output struct {
	Name            string             `json:"name"`
	Id              string             `json:"id"`
	InputReferences []*ActionReference `json:"inputs"`
	actionUid string
}

func NewActionOutput(name string, actionUid string) *Output {
	id := fmt.Sprintf("%s__Output:%s", actionUid, name)
	inputReferences := []*ActionReference{}
	return &Output{
		name,
		id,
		inputReferences,
		actionUid,
	}
}

func (ao *Output) AssignInputReference(ref *ActionReference) {
	ao.InputReferences = append(ao.InputReferences, ref)
}

func (ao *Output) UnmarshalJSON(data []byte) error {
	type alias Output
	var temp alias

	if _, err := marshalling.StrictDecode(data, &temp); err != nil {
		return err
	}

	if temp.Name != ao.Name {
		return fmt.Errorf("name mismatch: got '%s', expected '%s'", temp.Name, ao.Name)
	}

	*ao = Output(temp)
	return nil
}
