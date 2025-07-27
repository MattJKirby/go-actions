package action_test

import (
	"go-actions/ga/action"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"testing"
)

func TestNewActionInternals(t *testing.T) {
	
	reg := action.ActionRegistration[testActions.ActionValidEmpty]{Action: testActions.ActionValidEmpty{}}
	definition := action.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance(mockConfig, definition)
	
	internals := action.NewActionInternals(instance)

	assert.Equals(t, internals.Instance, instance)
}