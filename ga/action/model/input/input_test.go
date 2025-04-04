package input

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestAssignSourceReference(t *testing.T) {
	partial := &io.PartialActionReference{"refUid", "actionUid"}

	input := NewActionInput("name", "targetUid")
	input.AssignSourceReference(partial)

	stored, err := input.SourceReferences.Get(partial.ReferenceUid)
	assert.Equals(t, partial, stored)
	assert.Equals(t, nil, err)
}
