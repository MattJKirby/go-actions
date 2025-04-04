package output

import (
	"go-actions/ga/action/model/io"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestAssignTargetReference(t *testing.T) {
	partial := &io.PartialActionReference{"refUid", "actionUid"}
	output := NewActionOutput("name", "sourceUid")

	output.AssignTargetReference(partial)

	stored, err := output.TargetReferences.Get(partial.ReferenceUid)
	assert.Equals(t, nil, err)
	assert.Equals(t, partial, stored)
}
