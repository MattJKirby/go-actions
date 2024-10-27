package action

import (
	"go-actions/ga/action/instance"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewGoActionInternals(t *testing.T) {
	internals := NewGoActionInternals("test")
	expectedInstance := instance.NewActionInstance("test")
	expectedInstance.ActionUid = internals.Instance.ActionUid
	t.Run("test create new internals", func(t *testing.T) {
		asserts.Equals(t, expectedInstance, internals.Instance)
	})
}