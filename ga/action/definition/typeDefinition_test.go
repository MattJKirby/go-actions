package definition

import (
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"reflect"
	"testing"
)

func TestTypeDefinitionFromRegistration(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()

	expectedTypeName := "ActionValidEmpty"
	expectedTypePath := "go-actions/ga/testing/testActions/testActions.ActionValidEmpty"
	expectedType := reflect.TypeOf(ta.ActionValidEmpty{})
	expectedValue := reflect.ValueOf(&ta.ActionValidEmpty{})
	expectedCtor := reflect.ValueOf(reg.Constructor).Pointer()
	expectedCtorType := reflect.TypeOf(reg.Constructor)
	expectedPropsType := reflect.TypeOf(reg.DefaultProps)
	expectedPropsValue := reflect.ValueOf(reg.DefaultProps)

	defReg := TypeDefinitionFromRegistration(&reg)

	asserts.Equals(t, expectedTypeName, defReg.TypeName)
	asserts.Equals(t, expectedTypePath, defReg.TypePath)
	asserts.Equals(t, expectedType, defReg.ActionType)
	asserts.Equals(t, expectedValue, defReg.ActionValue)
	asserts.Equals(t, expectedCtor, defReg.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defReg.CtorType)
	asserts.Equals(t, expectedPropsType, defReg.PropsType)
	asserts.Equals(t, expectedPropsValue, defReg.PropsValue)
}
