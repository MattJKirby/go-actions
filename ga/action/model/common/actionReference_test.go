package common

import (
	"encoding/json"
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
	ref := NewActionReference(mockGlobalConfig, &sourceUid, &targetUid)

	expectedSourceRef := &ResourceReference{
		Uid: ref.Uid,
		Source: &sourceUid,
	}

	expecedTargetRef := &ResourceReference{
		Uid: ref.Uid,
		Target: &targetUid,
	}

	assert.Equals(t, expectedSourceRef, ref.GetSourceReference())
	assert.Equals(t, expecedTargetRef, ref.GetTargetReference())
}

func TestMarshalSourceAndTarget(t *testing.T){
	sourceUid := uid.NewUidBuilder().WithSubResource("source").Build()
	targetUid := uid.NewUidBuilder().WithSubResource("target").Build()
	ref := NewActionReference(mockGlobalConfig, &sourceUid, &targetUid)

	marshalledSource, srcErr := json.Marshal(ref.GetSourceReference())
	marshalledTarget, tgtErr := json.Marshal(ref.GetTargetReference())

	assert.Equals(t, `{"uid":"ga:core:ref:uid::","source":"ga:core:::source:"}`, string(marshalledSource))
	assert.Equals(t, `{"uid":"ga:core:ref:uid::","target":"ga:core:::target:"}`, string(marshalledTarget))
	assert.Equals(t, nil, srcErr)
	assert.Equals(t, nil, tgtErr)
}