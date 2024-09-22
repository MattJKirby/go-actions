package main

import (
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
)

func main(){

	ex,_ := ga.GetAction(examples.ExampleAction{})
	fmt.Println(ex)
}