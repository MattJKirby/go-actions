package examples

import "go-actions/ga/action"

func init(){
	action.NewAction(NewExampleAction)
}

type ExampleAction struct {
	abc string
}

func NewExampleAction() *ExampleAction {
	return &ExampleAction{
		abc: "adf",
	}
}