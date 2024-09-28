package examples

import (
	"go-actions/ga"
)

func init() {
	ga.DefineAction(NewExampleAction)
}

type ExampleAction struct {
	abc string
}

func NewExampleAction() *ExampleAction {
	return &ExampleAction{
		abc: "adf",
	}
}

func (ex *ExampleAction) Execute() {

}
