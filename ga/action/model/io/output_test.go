package io

import (
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidProvider: mockGenerator}

func TestAssignTargetReference(t *testing.T) {
	source := uid.NewUidBuilder().WithSubResource("source").Build()
	target := uid.NewUidBuilder().WithSubResource("target").Build()
	ref := NewActionReference(mockConfig, &source, &target)
	modelUid := uid.NewUidBuilder().Build()

	output := NewActionOutput(modelUid, "name")

	output.AssignTargetReference(ref.GetTargetReference())

	stored, err := output.TargetReferences.GetResource(ref.Uid.FullUid())
	assert.Equals(t, nil, err)
	assert.Equals(t, ref.GetTargetReference(), &stored)
}
