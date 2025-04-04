package output

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: ""}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestAssignTargetReference(t *testing.T) {
	ref := io.NewActionReference(mockGlobalConfig, "sourceUid", "targetUid")
	output := NewActionOutput("name", "sourceUid")

	output.AssignTargetReference(ref)

	stored, err := output.TargetReferences.Get(ref.ReferenceUid)
	assert.Equals(t, nil, err)
	assert.Equals(t, ref.GetTargetReference(), stored)
}
