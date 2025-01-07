package main

import (
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {
	definition, _ := ga.GetActionRegistration[examples.ExampleAction, examples.ExampleActionProps]()
	// // action, _ := ga.GetAction(examples.ExampleAction{})
	fmt.Println(definition)
	// fmt.Println(action)

	flow := ga.NewFlow()
	exa := examples.NewExampleAction(flow, nil)
	exa.Action.Execute()


}