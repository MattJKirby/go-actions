package output

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

func TestAssignTargetReference(t *testing.T) {
	source := uid.NewResourceUid(uid.WithSubResource("source"))
	target := uid.NewResourceUid(uid.WithSubResource("target"))
	ref := io.NewActionReference(mockConfig, source, target)
	modelUid := uid.NewResourceUid()

	output := NewActionOutput(modelUid, "name")

	output.AssignTargetReference(ref)

	stored, err := output.TargetReferences.GetResource(ref.Uid.GetUid())
	assert.Equals(t, nil, err)
	assert.Equals(t, ref, stored)
}
