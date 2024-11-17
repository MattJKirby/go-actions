package main

import (
	"encoding/json"
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main() {

	ex, _ := ga.GetActionDefinition(examples.ExampleAction{})
	fmt.Println(ex)

	action, _ := ga.GetAction(examples.ExampleAction{})

	fmt.Println(action)

	def := action.GetDef()

	def.Execute()
	def.IntegerParameter.SetValue(20)
	def.Execute()

	fmt.Println("--------")

	// gotIntParam := parameter.GetOrDefault("testInt", 0)

	inst := action.Instance
	marshalledInst, _ := json.Marshal(&inst)
	fmt.Println(string(marshalledInst))

}
