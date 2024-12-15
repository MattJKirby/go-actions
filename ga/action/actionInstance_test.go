package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("test")
	model := model.NewActionModel("test", &config.ModelConfig{})
	model.ActionUid = instance.ActionUid

	asserts.Equals(t, model, instance.ActionModel)
}
