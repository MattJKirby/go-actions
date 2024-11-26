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
	outputRef := reference.NewOutputReference("otherAction", "res")

	t.Run("assign output", func(t *testing.T) {
		input.AssignOutput(outputRef)
		asserts.Equals(t, outputRef, input.OutputReference)
	})
}

func TestMarshalling(t *testing.T) {
	input := newInput("i", "u")
	ref := reference.NewOutputReference("a", "o")
	marshalledRef, _ := json.Marshal(ref)

	t.Run("marshalling no output", func(t *testing.T) {
		marshalled, _ := json.Marshal(input)
		asserts.Equals(t, `{"name":"i","id":"u__Input:i","output":null}`, string(marshalled))
	})

	t.Run("marshalling with output", func(t *testing.T) {
		input.AssignOutput(ref)
		marshalled, _ := json.Marshal(input)
		expected := fmt.Sprintf(`{"name":"i","id":"u__Input:i","output":%s}`, string(marshalledRef))
		asserts.Equals(t, expected, string(marshalled))
	})
}

func TestUnmarshalling(t *testing.T) {
	input := newInput("i", "u")
	ref := reference.NewOutputReference("a", "o")
	input.AssignOutput(ref)
	marshalled, _ := json.Marshal(input)

	t.Run("test valid unmarshalling", func(t *testing.T) {
		newInput := newInput("i", "")
		json.Unmarshal(marshalled, newInput)

		asserts.Equals(t, input.Id, newInput.Id)
		asserts.Equals(t, input.Name, newInput.Name)
		asserts.Equals(t, input.OutputReference.ActionUid, newInput.OutputReference.ActionUid)
		asserts.Equals(t, input.OutputReference.ResourceName, newInput.OutputReference.ResourceName)
	})

	t.Run("test invalid unmarshalling", func(t *testing.T) {
		newInput := newInput("badName", "")
		err := json.Unmarshal(marshalled, newInput)
		if err == nil {
			t.Errorf("expected err but got %v", nil)
		}
	})
}
