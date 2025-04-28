package io

import (
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

type source struct {
	assignedTarget *PartialActionReference
}

func (ms *source) GetActionUid() string  { return "sAction" }
func (ms *source) GetResourceId() string { return "sProperty" }
func (ms *source) AssignTargetReference(pref *PartialActionReference) error {
	ms.assignedTarget = pref
	return nil
}

type target struct {
	assignedSource *PartialActionReference
}

func (ms *target) GetActionUid() string  { return "tAction" }
func (ms *target) GetResourceId() string { return "tProperty" }
func (ms *target) AssignSourceReference(pref *PartialActionReference) error {
	ms.assignedSource = pref
	return nil
}

func TestGetReferences(t *testing.T) {
	ref := NewActionReference(mockGlobalConfig, &source{}, &target{})
	expectedSource := &PartialActionReference{ReferenceUid: ref.referenceUid, ActionUid: "sAction"}
	expectedTarget := &PartialActionReference{ReferenceUid: ref.referenceUid, ActionUid: "tAction"}

	assert.Equals(t, expectedSource, ref.getSourceReference())
	assert.Equals(t, expectedTarget, ref.getTargetReference())
}

func TestAssignReferences(t *testing.T) {
	ms := &source{}
	mt := &target{}
	ref := NewActionReference(mockGlobalConfig, ms, mt)

	ref.AssignReferences()

	assert.Equals(t, ref.getTargetReference(), ms.assignedTarget)
	assert.Equals(t, ref.getSourceReference(), mt.assignedSource)
}
