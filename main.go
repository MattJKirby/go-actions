package main

import (
	"fmt"

	"go-actions/examples"
	"go-actions/ga"
	"go-actions/ga/flow"
)

func main() {
	// definition, _ := ga.GetActionRegistration[*examples.ExampleAction, examples.ExampleActionProps]()
	// action, _ := ga.GetAction[*examples.ExampleAction, examples.ExampleActionProps]()
	// fmt.Println(definition)
	// fmt.Println(action)

	fmt.Println(ga.GetRegisteredTypeDefinition[*examples.BasicAction]())

	// flow := ga.NewFlow()
	// a1, _ := flow.Definition.NewAction("BasicAction")
	// a2, _ := flow.Definition.NewAction("BasicAction")

	// err := flow.Definition.NewReference(a1.ActionOutput.Uid, a2.ActionInput.Uid)
	// fmt.Println(err)

	// flowDef, _ := json.Marshal(flow.Definition)
	// fmt.Println(string(flowDef))


	f := ga.NewFlow()
	ba, _ := flow.AddAction[*examples.BasicAction](f)
	

	fmt.Println(ba.Instance)

}
