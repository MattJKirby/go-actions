package main

import (
	"encoding/json"
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {
	definition, _ := ga.GetActionDefinition(examples.ExampleAction{})
	action, _ := ga.GetAction(examples.ExampleAction{})
	fmt.Println(definition)
	fmt.Println(action)

	action.GetDef()
	marshalledInstance, _ := json.Marshal(action.Instance)
	fmt.Println(string(marshalledInstance))
}
