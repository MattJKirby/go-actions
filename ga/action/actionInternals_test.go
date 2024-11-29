package action

import (
	"go-actions/ga/action/instance"
	"go-actions/ga/action/instance/config"
	"go-actions/ga/action/instance/io"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewGoActionInternals(t *testing.T) {
	internals := NewGoActionInternals("test")
	expectedInstance := instance.NewActionInstance("test", &config.InstanceConfig{})
	expectedInstance.ActionUid = internals.ActionUid
	expectedInstance.Inputs = io.NewStore[io.Input](internals.ActionUid)

	t.Run("test create new internals", func(t *testing.T) {
		asserts.Equals(t, expectedInstance, internals.ActionInstance)
	})
}