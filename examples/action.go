package examples

import (
	"fmt"
	"go-actions/ga"
	"go-actions/ga/action/instance/io"
	"go-actions/ga/action/instance/parameter"
	"go-actions/ga/action/internals"
)

func init() {
	ga.DefineAction(NewExampleAction)
}

type ExampleAction struct {
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *io.Input
}

func NewExampleAction(action internals.GoActionInternals) *ExampleAction {
	return &ExampleAction{
		IntegerParameter: parameter.GetOrDefault("intParam", 10)(action.Parameters),
		StringParameter:  parameter.GetOrDefault("strParam", "test")(action.Parameters),
		Input:            io.GetOrDefaultInput("test")(action.Inputs),
	}
}

func (ex ExampleAction) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}
