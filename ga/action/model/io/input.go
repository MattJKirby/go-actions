package io

import (
	"fmt"
	"go-actions/ga/utils/marshalling"
)

type Input struct {
	Name            string           `json:"name"`
	Id              string           `json:"id"`
	SourceReference *ActionReference `json:"source"`
	actionUid       string
	required        bool
}

func NewInput(name string, actionUid string, required bool) *Input {
	id := fmt.Sprintf("%s__Input:%s", actionUid, name)
	return &Input{
		name,
		id,
		nil,
		actionUid,
		required,
	}
}

func (i *Input) AssignSource(outputRef *ActionReference) {
	i.SourceReference = outputRef
}

func (i *Input) UnmarshalJSON(data []byte) error {
	type alias Input
	var temp alias

	if _, err := marshalling.StrictDecode(data, &temp); err != nil {
		return err
	}

	if temp.Name != i.Name {
		return fmt.Errorf("name mismatch: got '%s', expected '%s'", temp.Name, i.Name)
	}

	*i = Input(temp)
	return nil
}
