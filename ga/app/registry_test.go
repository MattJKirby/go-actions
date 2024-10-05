package app

import (
	"go-actions/ga/action"
	"go-actions/ga/cr"
	"reflect"
	"testing"
)

func TestAcceptDefinition(t *testing.T){
	defReg := NewActionDefinitionRegistry()
	def := action.ActionDefinition{}

	defReg.acceptDefinition(&def)
	abt := len(defReg.actionsByType)
	abn := len(defReg.actionsByName)

	if abt != 1 {
		t.Errorf("test actions by type: got: %d, expected: %d", abt, 1)
	}

	if abn != 1 {
		t.Errorf("test actions by name: got: %d, expected: %d", abt, 1)
	}
}

func TestGetDefinition(t *testing.T) {
	defReg := NewActionDefinitionRegistry()
	def := action.ActionDefinition{}
	defReg.acceptDefinition(&def)

	type expection struct {
		def    *action.ActionDefinition
		throws bool
	}

	tests := []cr.TestCase[reflect.Type, expection]{
		{Name: "existing def", Input: def.ActionType(), Expected: expection{def: &def, throws: false}},
		{Name: "not existing def", Input: reflect.TypeOf("err"), Expected: expection{def: nil, throws: true}},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, expection]) {
		storedDef, err := defReg.getDefinition(test.Input)

		if test.Expected.throws {
			if err == nil {
				t.Errorf("test %s: expected an error but got none", test.Name)
			}
			return
		} 
			
		if storedDef != test.Expected.def {
			t.Errorf("test %s: got %v, expected %v", test.Name, storedDef, test.Expected.def)
		}
	})
}
