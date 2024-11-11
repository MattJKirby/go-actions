package definition

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewDef(t *testing.T) {
	expectedTypeName := "myAction"
	expectedTypePath := "go-actions/ga/action/definition/definition.myAction"
	defCtor, _ := NewActionDefinition(newMyAction)

	t.Run("test def attrs", func(t *testing.T) {
		asserts.Equals(t, defCtor.Name, expectedTypeName)
		asserts.Equals(t, defCtor.TypePath, expectedTypePath)
		asserts.Equals(t, defCtor.ActionType, defCtor.ActionType)
		asserts.Equals(t, defCtor.CtorValue, defCtor.CtorValue)
	})
}
