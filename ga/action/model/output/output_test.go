package output

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testHelpers/actionTestHelpers"
	"testing"
)

var config = &actionTestHelpers.MockActionConfig{MockUid: ""}

func TestAssignTargetReference(t *testing.T) {
	ref := io.NewActionReference(config, "sourceUid", "targetUid")
	output := NewActionOutput("name", "sourceUid")

	output.AssignTargetReference(ref)

	asserts.Equals(t, ref, output.TargetReferences[ref.ReferenceUid])
}
