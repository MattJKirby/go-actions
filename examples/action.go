package examples

import (
	"go-actions/ga"
	"go-actions/ga/action"
)

func init() {
	ga.DefineAction(NewExampleAction)
}

type ExampleAction struct {
	Abc string
}

func NewExampleAction(internals action.GoActionInternals) *ExampleAction {
	return &ExampleAction{
		Abc: "adf",
	}
}

func (ex ExampleAction) Execute() {

}
