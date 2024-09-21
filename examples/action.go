package examples

import "go-actions/ga/action"

func init(){
	action.NewActionDefinition(NewExampleAction)
}

type ExampleAction struct {
	abc string
}

func NewExampleAction() *ExampleAction {
	return &ExampleAction{
		abc: "adf",
	}
}