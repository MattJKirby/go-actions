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
	expectedSourceRef := NewReference("b", "o")

	target1 := NewInput("i", "a", false)
	target2 := NewInput("i", "c", false)

	expectedTargetRefs := []*ActionReference{
		NewReference("a", "i"),
		NewReference("c", "i"),
	}

	AssignReferences(source, []*Input{target1, target2})

	asserts.Equals(t, expectedTargetRefs, source.TargetReferences)
	asserts.Equals(t, expectedSourceRef, target1.SourceReference)
	asserts.Equals(t, expectedSourceRef, target2.SourceReference)
}
