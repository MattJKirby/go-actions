package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action/model/io/reference"
)

type Input struct {
	Name            string                     `json:"name"`
	Id              string                     `json:"id"`
	OutputReference *reference.OutputReference `json:"output"`
}

func newInput(name string, actionUid string) *Input {
	id := fmt.Sprintf("%s__Input:%s", actionUid, name)
	return &Input{
		name,
		id,
		nil,
	}
}

func (i *Input) AssignOutput(ref *reference.OutputReference) {
	i.OutputReference = ref
}

func (i *Input) UnmarshalJSON(data []byte) error {
	type alias Input
	var temp alias

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.Name != i.Name {
		return fmt.Errorf("name mismatch: got '%s', expected '%s'", temp.Name, i.Name)
	}

	*i = Input(temp)
	return nil
}

func GetOrDefaultInput(name string) func(*IOStore[Input]) *Input {
	return func(s *IOStore[Input]) *Input {
		return s.GetOrDefault(name, newInput)
	}
}
