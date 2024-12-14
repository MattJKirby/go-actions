package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/resources"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("test")
	model := model.NewActionModel("test", &config.ModelConfig{})
	model.ActionUid = instance.ActionUid
	model.Inputs = resources.NewResourceStore[io.Input]()
	model.Outputs = resources.NewResourceStore[io.Output]()

	asserts.Equals(t, model, instance.ActionModel)
}
