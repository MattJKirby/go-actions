package common

import (
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestResourceReference(t *testing.T) {
	sourceUid := uid.ResourceUid{}
	targetUid := uid.ResourceUid{}
	ref := NewActionReference(mockGlobalConfig, sourceUid, targetUid)

	expectedSourceRef := &ResourceReference{
		Uid: ref.Uid,
		Resource: sourceUid,
	}

	expecedTargetRef := &ResourceReference{
		Uid: ref.Uid,
		Resource: targetUid,
	}

	assert.Equals(t, expectedSourceRef, ref.GetSourceReference())
	assert.Equals(t, expecedTargetRef, ref.GetTargetReference())
}