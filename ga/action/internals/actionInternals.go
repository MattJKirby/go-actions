package internals

import (
	"go-actions/ga/action/instance"
	"go-actions/ga/action/instance/config"
)

type GoActionInternals struct {
	*instance.ActionInstance
}

func NewGoActionInternals(actionName string) GoActionInternals {
	return GoActionInternals{
		ActionInstance: instance.NewActionInstance(actionName, config.NewInstanceConfig()),
	}
}