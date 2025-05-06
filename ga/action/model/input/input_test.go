package input

import (
	"go-actions/ga/action/model/common"
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestAssignSourceReference(t *testing.T) {

	source := uid.NewUidBuilder().WithSubResource("source").Build()
	target := uid.NewUidBuilder().WithSubResource("target").Build()
	ref := common.NewActionReference(mockConfig, source, target)
	modelUid := uid.NewUidBuilder().Build()

	input := NewActionInput(modelUid, "name")
	input.AssignSourceReference(ref.GetSourceReference())

	stored, err := input.SourceReferences.GetResource(ref.Uid.FullUid())
	assert.Equals(t, ref.GetSourceReference(), stored)
	assert.Equals(t, nil, err)
}
