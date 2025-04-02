package io

import (
	"go-actions/ga/testing/assert"
	"go-actions/ga/testing/testHelpers/actionTestHelpers"
	"testing"
)

var config = &actionTestHelpers.MockActionConfig{MockUid: ""}

func TestGetSourceReference(t *testing.T) {
	ref := NewActionReference(config, "sUid", "tUid")
	expected := &PartialActionReference{ReferenceUid: ref.ReferenceUid, ActionUid: "sUid"}
	assert.Equals(t, expected, ref.GetSourceReference())
}

func TestGetTargetReference(t *testing.T) {
	ref := NewActionReference(config, "sUid", "tUid")
	expected := &PartialActionReference{ReferenceUid: ref.ReferenceUid, ActionUid: "tUid"}
	assert.Equals(t, expected, ref.GetTargetReference())
}
