package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
)

type ActionInstance struct {
	Model *model.ActionModel
}

func NewActionInstance(actionName string) ActionInstance {
	return ActionInstance{
		Model: model.NewActionModel(actionName, config.NewModelConfig()),
	}
}
