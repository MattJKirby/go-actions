package main

import (
	"encoding/json"
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {
	definition, _ := ga.GetActionRegistration[*examples.ExampleAction, examples.ExampleActionProps]()
	action, _ := ga.GetAction[*examples.ExampleAction, examples.ExampleActionProps]()
	fmt.Println(definition)
	fmt.Println(action)

	flow := ga.NewFlow()
	ex1, _ := examples.NewExampleAction(flow, &examples.ExampleActionProps{
		IntProp: 999,
		StrProp: "aaa",
	})

	ex1.Execute()

	ex2, _ := examples.NewExampleAction(flow, nil)
	ex2.Execute()

	val, _ := json.Marshal(action.Instance)
	fmt.Println(string(val))

}
