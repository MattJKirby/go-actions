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

	stored, err := output.TargetReferences.Get(ref.ReferenceUid)
	asserts.Equals(t, nil, err)
	asserts.Equals(t, ref, stored)
}
