package main

import (
	"encoding/json"
	"fmt"
	"go-actions/examples"
	"go-actions/ga"
	"go-actions/ga/action/instance/io/reference"
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

	fmt.Println(def.Input)
	outRef := reference.NewOutputReference("someAction", "someResource")

	def.Input.AssignOutput(outRef)

	fmt.Println("--------")

	// gotIntParam := parameter.GetOrDefault("testInt", 0)

	inst := action.Instance
	marshalledInst, _ := json.Marshal(&inst)
	fmt.Println(string(marshalledInst))
	fmt.Println("--------")

	mInst := `{"name":"ExampleAction","uid":"Action:ExampleAction:dad7805d-3cd1-4622-912e-c560398e9af1","parameters":{"intParam":{"name":"intParam","value":200},"strParam":{"name":"strParam","value":"HELL)"}},"inputs":{"test":{"name":"test","id":"Action:ExampleAction:dad7805d-3cd1-4622-912e-c560398e9af1__Input:test"}}}`
	json.Unmarshal([]byte(mInst), inst)

	def.Execute()

}
