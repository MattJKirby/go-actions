package definition

import (
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"testing"
)

func TestNewDef(t *testing.T) {
	expectedTypeName := "EmptyActionValid"
	expectedTypePath := "go-actions/ga/testing/testActions/testActions.EmptyActionValid"
	reg := ta.GenerateEmptyActionValidRegistration()

	defCtor, _ := NewActionDefinition(&reg)

	asserts.Equals(t, defCtor.Name, expectedTypeName)
	asserts.Equals(t, defCtor.TypePath, expectedTypePath)
	// asserts.Equals(t, defCtor.ActionType, defCtor.ActionType)
	// asserts.Equals(t, defCtor.CtorValue, defCtor.CtorValue)
}
