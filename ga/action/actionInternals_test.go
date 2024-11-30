package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
	"go-actions/ga/action/model/io"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewGoActionInternals(t *testing.T) {
	internals := NewGoActionInternals("test")
	expectedInstance := model.NewActionInstance("test", &config.InstanceConfig{})
	expectedInstance.ActionUid = internals.ActionUid
	expectedInstance.Inputs = io.NewStore[io.Input](internals.ActionUid)

	t.Run("test create new internals", func(t *testing.T) {
		asserts.Equals(t, expectedInstance, internals.ActionInstance)
	})
}
