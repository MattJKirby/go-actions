package io

import (
	"fmt"
	"go-actions/ga/utils/marshalling"
)

type Output struct {
	Name             string             `json:"name"`
	Id               string             `json:"id"`
	TargetReferences []*ActionReference `json:"targets"`
	actionUid        string
}

func NewActionOutput(name string, actionUid string) *Output {
	id := fmt.Sprintf("%s__Output:%s", actionUid, name)
	TargetReferences := []*ActionReference{}
	return &Output{
		name,
		id,
		TargetReferences,
		actionUid,
	}
}

func (ao *Output) AssignTarget(inputRef *ActionReference) {
	ao.TargetReferences = append(ao.TargetReferences, inputRef)
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
