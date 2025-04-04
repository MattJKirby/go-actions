package output

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers/actionTestHelpers"
	"testing"
)

var config = &actionTestHelpers.MockActionConfig{MockUid: ""}

func TestAssignTargetReference(t *testing.T) {
	ref := io.NewActionReference(config, "sourceUid", "targetUid")
	output := NewActionOutput("name", "sourceUid")

	output.AssignTargetReference(ref)

	stored, err := output.TargetReferences.Get(ref.ReferenceUid)
	assert.Equals(t, nil, err)
	assert.Equals(t, ref.GetTargetReference(), stored)
}
