package input

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestAssignSourceReference(t *testing.T) {
	ref := io.NewActionReference(mockGlobalConfig, "sourceUid", "targetUid")
	input := NewActionInput("name", "targetUid")

	input.AssignSourceReference(ref)

	stored, err := input.SourceReferences.Get(ref.ReferenceUid)
	assert.Equals(t, ref.GetSourceReference(), stored)
	assert.Equals(t, nil, err)
}
