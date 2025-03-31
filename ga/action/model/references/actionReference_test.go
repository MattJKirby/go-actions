package references

import (
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testHelpers/actionTestHelpers"
	"testing"
)

var config = &actionTestHelpers.MockActionConfig{MockUid: ""}

func TestGetSourceReference(t *testing.T) {
	ref := NewActionReference(config, "sUid", "tUid")
	expected := &PartialActionReference{ReferenceUid: ref.ReferenceUid, ActionUid: "sUid"}
	asserts.Equals(t, expected, ref.GetSourceReference())
}

func TestGetTargetReference(t *testing.T) {
	ref := NewActionReference(config, "sUid", "tUid")
	expected := &PartialActionReference{ReferenceUid: ref.ReferenceUid, ActionUid: "tUid"}
	asserts.Equals(t, expected, ref.GetTargetReference())
}
