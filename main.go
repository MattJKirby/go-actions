package main

import (
	"fmt"
	"go-actions/examples"
	"go-actions/ga/action"
)

func main(){

	testFn := action.NewActionDefinition(examples.NewExampleAction)
	testStruct := action.NewActionDefinition(&examples.ExampleAction{})

	fmt.Println(testFn.Name(), testFn.TypeName(), testFn.Type(), testFn.Value())
	fmt.Println(testStruct.Name(), testStruct.TypeName(), testStruct.Type(), testStruct.Value())
}