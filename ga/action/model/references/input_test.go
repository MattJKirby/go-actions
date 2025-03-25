package references

import (
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testHelpers/actionTestHelpers"
	"testing"
)

var config = &actionTestHelpers.MockActionConfig{MockUid: ""}

func TestAssignSourceReference(t *testing.T) {
	ref := NewActionReference(config, "sourceUid", "targetUid")
	input := NewActionInput("name", "targetUid")

	input.AssignSourceReference(ref)

	stored, err := input.SourceReferences.Get(ref.ReferenceUid)
	asserts.Equals(t, ref, stored)
	asserts.Equals(t, nil, err)
}