package instance

import (
	"testing"
)

func TestNewActionInstance(t *testing.T) {
	typeName := "someName"
	instance := NewActionInstance(typeName)

	t.Run("test new instance", func(t *testing.T) {
		if instance.ActionName != typeName {
			t.Errorf("invlaid typename: expected %s, got %s", typeName, instance.ActionName)
		}

		if instance.ActionUid == "" {
			t.Errorf("expected non-empty action uid")
		}

		if instance.Parameters == nil {
			t.Errorf("invalid parameter store: expected parameter store, got %v", nil)
		}

		if instance.Inputs == nil {
			t.Errorf("invalid reference store: expected reference store got %v", nil)
		}
	})
}
