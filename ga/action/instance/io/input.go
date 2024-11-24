package io

import "fmt"

type Input struct {
	name string
	id string
}

func newInput(name string, actionUid string) *Input {
	id := fmt.Sprintf("%s__Input:%s", actionUid, name)
	return &Input{
		name,
		id,
	}
}

func (i *Input) Name() string {
	return i.name
}

func (i *Input) Id() string {
	return i.id
}