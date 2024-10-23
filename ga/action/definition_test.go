package action

import (
	"go-actions/ga/cr/asserts"
	"testing"
)



func TestNewDef(t *testing.T){
	expectedTypeName := "myAction"
	expectedTypePath := "go-actions/ga/action/action.myAction"
	defCtor, _ := NewActionDefinition[myAction](newMyAction)
	defStruct, _ := NewActionDefinition[myAction](myAction{})
	defInvalid, err := NewActionDefinition[myAction](10)

	t.Run("test def attrs", func(t *testing.T) {
		asserts.Equals(t, defCtor.Name(), expectedTypeName)
		asserts.Equals(t, defStruct.Name(), expectedTypeName)
		asserts.Equals(t, defCtor.TypeName(), expectedTypePath)
		asserts.Equals(t, defStruct.TypeName(), expectedTypePath)
		
		if defInvalid != nil || err == nil {
			t.Errorf("expected error from invalid definition")
		}
	})
}