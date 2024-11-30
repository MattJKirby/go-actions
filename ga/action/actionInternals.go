package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/config"
)

type GoActionInternals struct {
	*model.ActionInstance
}

func NewGoActionInternals(actionName string) GoActionInternals {
	return GoActionInternals{
		ActionInstance: model.NewActionInstance(actionName, config.NewInstanceConfig()),
	}
}
