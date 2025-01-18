package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewInput(t *testing.T) {
	input := NewInput("name", "actionUid", false)

	asserts.Equals(t, "name", input.Name)
	asserts.Equals(t, "actionUid__Input:name", input.Id)
}

func TestAssignOutput(t *testing.T) {
	input := NewInput("name", "actionUid", false)
	outputRef := NewReference("otherAction", "res")

	input.AssignSource(outputRef)
	asserts.Equals(t, outputRef, input.SourceReference)
}

func TestMarshalling(t *testing.T) {
	input := NewInput("i", "u", false)
	ref := NewReference("a", "o")
	marshalledRef, _ := json.Marshal(ref)

	t.Run("marshalling no output", func(t *testing.T) {
		marshalled, _ := json.Marshal(input)
		asserts.Equals(t, `{"name":"i","id":"u__Input:i","source":null}`, string(marshalled))
	})

	t.Run("marshalling with output", func(t *testing.T) {
		input.AssignSource(ref)
		marshalled, _ := json.Marshal(input)
		expected := fmt.Sprintf(`{"name":"i","id":"u__Input:i","source":%s}`, string(marshalledRef))
		asserts.Equals(t, expected, string(marshalled))
	})
}

func TestUnmarshalling(t *testing.T) {
	input := NewInput("i", "u", false)
	ref := NewReference("a", "o")
	input.AssignSource(ref)
	marshalled, _ := json.Marshal(input)

	t.Run("test valid unmarshalling", func(t *testing.T) {
		newInput := NewInput("i", "", false)
		json.Unmarshal(marshalled, newInput)

		asserts.Equals(t, input.Id, newInput.Id)
		asserts.Equals(t, input.Name, newInput.Name)
		asserts.Equals(t, input.SourceReference.ActionUid, newInput.SourceReference.ActionUid)
		asserts.Equals(t, input.SourceReference.ResourceName, newInput.SourceReference.ResourceName)
	})

	t.Run("test invalid unmarshalling", func(t *testing.T) {
		newInput := NewInput("badName", "", false)
		err := json.Unmarshal(marshalled, newInput)
		if err == nil {
			t.Errorf("expected err but got %v", nil)
		}
	})
}
