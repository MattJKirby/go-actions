package definition

import (
	"go-actions/ga/action"
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

func TestNewDefaultProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	newProps := defReg.NewDefaultProps()
	typeAssertedProps, ok := (newProps).(*ta.ActionValidProps)

	asserts.Equals(t, true, ok)
	asserts.Equals(t, reg.DefaultProps, typeAssertedProps)
}

func TestNewConstructor(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	expectedInst := action.NewActionInstance("expected")
	expectedAction := reg.Constructor(expectedInst, ta.ActionValidProps{Param1: "somePropValue"})

	testInst := action.NewActionInstance("test")
	testCtor := defReg.NewConstructor()
	testAction := testCtor(testInst, ta.ActionValidProps{Param1: "somePropValue"})

	typedTestAction, ok := (testAction).(*ta.ActionValid)

	asserts.Equals(t, true, ok)
	asserts.Equals(t, expectedAction, typedTestAction)
	asserts.Equals(t, expectedInst.Model.Parameters, testInst.Model.Parameters)

}
