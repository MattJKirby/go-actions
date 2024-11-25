package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action/instance/io/reference"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewInput(t *testing.T) {
	input := newInput("name", "actionUid")

	z, _ := json.Marshal(input)
	fmt.Println(string(z))

	t.Run("test new input", func(t *testing.T) {
		asserts.Equals(t, "name", input.Name)
		asserts.Equals(t, "actionUid__Input:name", input.Id)
	})
}

func TestGetOrDefault(t *testing.T) {
	store := NewStore[Input]("uid")

	t.Run("test default", func(t *testing.T) {
		expected := newInput("name", "uid")
		input := GetOrDefaultInput("name")(store)

		asserts.Equals(t, expected, input)

	})
}

func TestAssignOutput(t *testing.T) {
	input := newInput("name", "actionUid")
	outputRef := reference.NewReference("otherAction", "res", "output")

	t.Run("assign output", func(t *testing.T) {
		input.AssignOutput(outputRef)
		asserts.Equals(t, outputRef, input.OutputReference)
	})
}

func TestMarshalling(t *testing.T) {
	input := newInput("name", "actionUid")

	t.Run("marshalling no output", func(t *testing.T) {
		marshalled, _ := json.Marshal(input)
		asserts.Equals(t, `{"name":"name","id":"actionUid__Input:name","outputRef":null}`, string(marshalled))
	})

	t.Run("marshalling with output", func(t *testing.T) {
		outputRef := reference.NewReference("otherAction", "res", "output")
		marshalledRef, _ := json.Marshal(outputRef)
		input.AssignOutput(outputRef)
		marshalled, _ := json.Marshal(input)
		expected := fmt.Sprintf( `{"name":"name","id":"actionUid__Input:name","outputRef":%s}`, string(marshalledRef))
		asserts.Equals(t,expected, string(marshalled))
	})
}
