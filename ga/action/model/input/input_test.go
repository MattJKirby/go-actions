package input

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var config = &testHelpers.MockUidGenerator{MockUid: ""}

func TestAssignSourceReference(t *testing.T) {
	ref := io.NewActionReference(config, "sourceUid", "targetUid")
	input := NewActionInput("name", "targetUid")

	input.AssignSourceReference(ref)

	stored, err := input.SourceReferences.Get(ref.ReferenceUid)
	assert.Equals(t, ref.GetSourceReference(), stored)
	assert.Equals(t, nil, err)
}
