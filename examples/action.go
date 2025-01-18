package examples

import (
	"fmt"
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/flow"
)

func init() {
	ga.RegisterAction(&action.GoActionRegistration[ExampleAction, ExampleActionProps]{
		Constructor: newExampleAction,
		DefaultProps: &ExampleActionProps{
			IntProp: 10,
			StrProp: "someString",
		},
	})
}

type ExampleActionProps struct {
	IntProp int
	StrProp string
	Source   *io.Output
	Targets []*io.Input
}

type ExampleAction struct {
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *io.Input
	Output           *io.Output
}

func newExampleAction(instance *action.ActionInstance, props ExampleActionProps) *ExampleAction {
	return &ExampleAction{
		IntegerParameter: action.Parameter(instance, "intParam", props.IntProp),
		StringParameter:  action.Parameter(instance, "strParam", props.StrProp),
		Input:            action.Input(instance, "input", true),
		Output:           action.Output(instance, "output"),
	}
}

func (ex ExampleAction) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}

func NewExampleAction(flow *flow.Flow, props *ExampleActionProps) *executable.Action[ExampleAction, ExampleActionProps] {
	return ga.NewFlowAction[ExampleAction](flow, props)
}
