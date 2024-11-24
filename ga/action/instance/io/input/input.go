package input

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action/instance/io"
)

type Input struct {
	name string 
	id   string 
}

type MarshalledInput struct {
	Name string `json:"name"`
}

func newInput(name string, actionUid string) *Input {
	id := fmt.Sprintf("%s__Input:%s", actionUid, name)
	return &Input{
		name,
		id,
	}
}

func GetOrDefaultInput(name string) func(*io.Store[Input]) *Input {
	return func(s *io.Store[Input]) *Input {
		return s.GetOrDefault(name, newInput)
	}
}

func (i *Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(&MarshalledInput{
		Name: "vvv",
	})
}
