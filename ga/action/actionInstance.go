package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
)

type ActionInstance struct {
	*model.ActionModel
}

func NewActionInstance(actionName string) ActionInstance {
	return ActionInstance{
		ActionModel: model.NewActionModel(actionName, config.NewModelConfig()),
	}
}
