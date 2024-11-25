package io

import (
	"fmt"
	"go-actions/ga/action/instance/io/reference"
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

func GetOrDefaultInput(name string) func(*Store[Input]) *Input {
	return func(s *Store[Input]) *Input {
		return s.GetOrDefault(name, newInput)
	}
}
