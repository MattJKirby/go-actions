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
  source := NewActionOutput("o", "b")
  sourceRef := NewReference("b", "o")

	target1 := NewInput("i", "a", false)
	targetRef1 := NewReference("a", "i")

  target2 := NewInput("i", "c", false)
	targetRef2 := NewReference("c", "i")

	AssignReferences(source, []*Input{target1, target2})

	asserts.Equals(t, []*ActionReference{targetRef1, targetRef2}, source.InputReferences)
	asserts.Equals(t, sourceRef, target1.OutputReference)
  asserts.Equals(t, sourceRef, target2.OutputReference)
}
