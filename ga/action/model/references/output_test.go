package references

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestAssignTargetReference(t *testing.T) {
	ref := NewActionReference(config, "sourceUid", "targetUid")
	output := NewActionOutput("name", "sourceUid")

	output.AssignTargetReference(ref)

	asserts.Equals(t, ref, output.TargetReferences[ref.ReferenceUid])
}