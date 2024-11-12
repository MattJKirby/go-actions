package main

import (
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
}
