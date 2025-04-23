package examples

import (
	"fmt"
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
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
	Source  *output.ActionOutput
	Targets []*input.ActionInput
}

type ExampleAction struct {
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *input.ActionInput
	Output           *output.ActionOutput
}

func newExampleAction(instance *action.ActionInstance, props ExampleActionProps) *ExampleAction {
	return &ExampleAction{
		IntegerParameter: model.Parameter(instance.Model, "intParam", props.IntProp),
		StringParameter:  model.Parameter(instance.Model, "strParam", props.StrProp),
		Input:            model.Input(instance.Model, "input", true, props.Source),
		Output:           model.Output(instance.Model, "output", props.Targets),
	}
}

func (ex ExampleAction) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}
