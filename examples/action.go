package examples

import (
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/instance/parameter"
)

func init() {
	ga.DefineAction(NewExampleAction)
}

type ExampleAction struct {
	Abc string
	intParam *parameter.ActionParameter[int]
}

func NewExampleAction(action action.GoActionInternals) *ExampleAction {
	return &ExampleAction{
		Abc: "adf",
		intParam: parameter.GetOrDefault("int", 10)(action.Parameters),

	}
}

func (ex ExampleAction) Execute() {

}
