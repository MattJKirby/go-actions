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

	action.GetDef()
}
