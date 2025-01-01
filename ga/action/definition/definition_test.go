package definition

import (
	"go-actions/ga/cr/asserts"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewDef(t *testing.T) {
	expectedTypeName := "TestActionValid"
	expectedTypePath := "go-actions/ga/utils/testHelpers/testHelpers.TestActionValid"
	defCtor, _ := NewActionDefinition(th.GetEmptyConstructor[th.TestActionValid]())

	asserts.Equals(t, defCtor.Name, expectedTypeName)
	asserts.Equals(t, defCtor.TypePath, expectedTypePath)
	asserts.Equals(t, defCtor.ActionType, defCtor.ActionType)
	asserts.Equals(t, defCtor.CtorValue, defCtor.CtorValue)
}
