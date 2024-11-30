package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
)

type GoActionInternals struct {
	*model.ModelInstance
}

func NewGoActionInternals(actionName string) GoActionInternals {
	return GoActionInternals{
		ModelInstance: model.NewModelInstance(actionName, config.NewInstanceConfig()),
	}
}
