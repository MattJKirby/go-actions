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
	expectedInstance := model.NewModelInstance("test", &config.InstanceConfig{})
	expectedInstance.ActionUid = internals.ActionUid
	expectedInstance.Inputs = io.NewIOStore[io.Input](internals.ActionUid)

	asserts.Equals(t, expectedInstance, internals.ModelInstance)
}
