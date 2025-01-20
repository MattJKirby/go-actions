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

func TestAssignNilReferences(t *testing.T) {
	testcases := []struct {
		name    string
		source  *Output
		targets []*Input
	}{
		{name: "nil source", source: nil, targets: []*Input{NewInput("i", "a", false)}},
		{name: "nil targets", source: NewActionOutput("o", "b"), targets: []*Input{nil}},
	}

	for _, test := range testcases {
		t.Helper()
		t.Run(test.name, func(t *testing.T) {
			AssignReferences(test.source, test.targets)
			if test.source != nil {
				asserts.Equals(t, len(test.source.TargetReferences), 0)
			}

			for _, tar := range test.targets {
				if tar != nil {
					asserts.Equals(t, tar.SourceReference, nil)
				}
			}
		})
	}
}
