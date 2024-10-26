package action

import "testing"

func TestNewActionInstance(t *testing.T){
	typeName := "someName"
	def := ActionDefinition{Name: typeName}
	instance := NewActionInstance(&def)

	t.Run("test new instance", func(t *testing.T){
		if instance.ActionName != typeName {
			t.Errorf("invlaid typename: expected %s, got %s", typeName, instance.ActionName)
		}

		if instance.ActionUid == "" {
			t.Errorf("expected non-empty action uid")
		}
	})
}