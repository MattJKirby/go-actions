package input

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestAssignSourceReference(t *testing.T) {
	
	source := uid.NewResourceUid(uid.WithSubResource("source"))
	target := uid.NewResourceUid(uid.WithSubResource("target"))
	ref := io.NewActionReference(mockConfig, source, target)
	modelUid := uid.NewResourceUid()

	input := NewActionInput(modelUid, "name")
	input.AssignSourceReference(ref)

	stored, err := input.SourceReferences.Get(ref.Uid.GetString())
	assert.Equals(t, ref, stored)
	assert.Equals(t, nil, err)
}
