package definition

import (
	"go-actions/ga/cr/asserts"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewDef(t *testing.T) {
	expectedTypeName := "ActionValid"
	expectedTypePath := "go-actions/ga/utils/testHelpers/testHelpers.ActionValid"
	defCtor, _ := NewActionDefinition(th.GetEmptyConstructor[th.ActionValid, th.ActionValidProps]())

	asserts.Equals(t, defCtor.Name, expectedTypeName)
	asserts.Equals(t, defCtor.TypePath, expectedTypePath)
	asserts.Equals(t, defCtor.ActionType, defCtor.ActionType)
	asserts.Equals(t, defCtor.CtorValue, defCtor.CtorValue)
}
