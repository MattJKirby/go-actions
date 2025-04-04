package io

import (
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestGetSourceReference(t *testing.T) {
	ref := NewActionReference(mockGlobalConfig, "sUid", "tUid")
	expected := &PartialActionReference{ReferenceUid: ref.ReferenceUid, ActionUid: "sUid"}
	assert.Equals(t, expected, ref.GetSourceReference())
}

func TestGetTargetReference(t *testing.T) {
	ref := NewActionReference(mockGlobalConfig, "sUid", "tUid")
	expected := &PartialActionReference{ReferenceUid: ref.ReferenceUid, ActionUid: "tUid"}
	assert.Equals(t, expected, ref.GetTargetReference())
}
