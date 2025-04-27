package main

import (
	"encoding/json"
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {
	// definition, _ := ga.GetActionRegistration[*examples.ExampleAction, examples.ExampleActionProps]()
	// action, _ := ga.GetAction[*examples.ExampleAction, examples.ExampleActionProps]()
	// fmt.Println(definition)
	// fmt.Println(action)

	flow := ga.NewFlow()
	ex1, _ := examples.NewExampleAction(flow)
	fmt.Println(ex1)

	ex2, _ := examples.NewBasicAction(flow)

	inst, _ := json.Marshal(ex2.Instance)

	fmt.Println(string(inst))

}
