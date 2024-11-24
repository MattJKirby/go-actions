package input

import (
	"fmt"
	"go-actions/ga/action/instance/io"
)

type Input struct {
	name string
	id   string
}

func newInput(name string, actionUid string) *Input {
	id := fmt.Sprintf("%s__Input:%s", actionUid, name)
	return &Input{
		name,
		id,
	}
}

func (i Input) Name() string {
	return i.name
}

func (i Input) Id() string {
	return i.id
}

func GetOrDefaultInput(name string) func(*io.Store[Input]) *Input {
	return func(s *io.Store[Input]) *Input {
		return s.GetOrDefaultResource(name, newInput)
	}
}
