package io

import (
	"fmt"
)

type Input struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func newInput(name string, actionUid string) *Input {
	id := fmt.Sprintf("%s__Input:%s", actionUid, name)
	return &Input{
		name,
		id,
	}
}

func GetOrDefaultInput(name string) func(*Store[Input]) *Input {
	return func(s *Store[Input]) *Input {
		return s.GetOrDefault(name, newInput)
	}
}
