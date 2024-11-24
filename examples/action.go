package examples

import (
	"fmt"
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/instance/io/input"
	"go-actions/ga/action/instance/parameter"
)

func init() {
	ga.DefineAction(NewExampleAction)
}

type ExampleAction struct {
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *input.Input
}

func NewExampleAction(action action.GoActionInternals) *ExampleAction {
	return &ExampleAction{
		IntegerParameter: parameter.GetOrDefault("intParam", 10)(action.Parameters),
		StringParameter:  parameter.GetOrDefault("strParam", "test")(action.Parameters),
		Input:            input.GetOrDefaultInput("test")(action.Inputs),
	}
}

func (ex ExampleAction) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}
