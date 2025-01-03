package main

import (
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {
	definition, _ := ga.GetActionRegistration(examples.ExampleAction{})
	action, _ := ga.GetAction(examples.ExampleAction{})
	fmt.Println(definition)
	fmt.Println(action)

	// action.GetDef()
	// marshalledInstance, _ := json.Marshal(action.Instance)
	// fmt.Println(string(marshalledInstance))

	flow := ga.NewFlow()
	exa := ga.ActionFunction(flow, examples.ExampleAction{}, examples.ExampleActionProps{})

	exa.GetDef().Execute()

}
