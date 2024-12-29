package main

import (
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {
	definition, _ := ga.GetActionDefinition(examples.ExampleActionConstruct{})
	action, _ := ga.GetAction(examples.ExampleActionConstruct{})
	fmt.Println(definition)
	fmt.Println(action)

	// action.GetDef()
	// marshalledInstance, _ := json.Marshal(action.Instance)
	// fmt.Println(string(marshalledInstance))

	flow := ga.NewFlow()
	exa := ga.NewAction(flow, examples.ExampleActionConstruct{})
	
	
	exa.GetDef().Execute()

}
