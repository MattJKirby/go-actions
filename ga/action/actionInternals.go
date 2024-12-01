package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
)

type GoActionInternals struct {
	*model.ActionModel
}

func NewGoActionInternals(actionName string) GoActionInternals {
	return GoActionInternals{
		ActionModel: model.NewActionModel(actionName, config.NewModelConfig()),
	}
}
