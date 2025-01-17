package io

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestMarshalReference(t *testing.T) {
	outputRef := NewReference("a", "o")
	marshalledOutput, _ := json.Marshal(outputRef)
	asserts.Equals(t, `{"actionUid":"a","resourceName":"o"}`, string(marshalledOutput))
}

func TestUnmarshalOutputReference(t *testing.T) {
	marshalledOutput := []byte(`{"actionUid":"a","resourceName":"o"}`)
	outputRef := NewReference("", "")

	json.Unmarshal(marshalledOutput, outputRef)
	asserts.Equals(t, "a", outputRef.ActionUid)
	asserts.Equals(t, "o", outputRef.ResourceName)
}

func TestAssignReferences(t *testing.T) {
	input := NewInput("i", "a", false)
	output := NewActionOutput("o", "b")
	inputRef := NewReference("a", "i")
	outputRef := NewReference("b", "o")

	AssignReferences(input, output)

	asserts.Equals(t, outputRef, input.OutputReference)
	asserts.Equals(t, inputRef, output.InputReferences[0])
}
