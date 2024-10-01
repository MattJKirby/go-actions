package examples

import (
	"go-actions/ga"
)

func init() {
	ga.DefineAction(NewExampleAction)
}

type ExampleAction struct {
	Abc string
}

func NewExampleAction() *ExampleAction {
	return &ExampleAction{
		Abc: "adf",
	}
}

func (ex *ExampleAction) Execute() {

}
