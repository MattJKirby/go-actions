package main

import (
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {
	definition, _ := ga.GetActionRegistration[examples.ExampleAction, examples.ExampleActionProps]()
	action, _ := ga.GetAction[examples.ExampleAction, examples.ExampleActionProps]()
	fmt.Println(definition)
	fmt.Println(action)

	flow := ga.NewFlow()
	ex1 := examples.NewExampleAction(flow, &examples.ExampleActionProps{
		IntProp: 999,
	})

	
	ex1.Action.Execute()

	ex2 := examples.NewExampleAction(flow, nil)
	ex2.Action.Execute()

}
