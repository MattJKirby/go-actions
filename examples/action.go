package examples

import (
	"go-actions/ga"
	"go-actions/ga/action"
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

func Test() action.ActionFunction {
	return NewExampleAction()
}