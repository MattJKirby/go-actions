package action

import "go-actions/ga/action/instance"

type GoActionConstructor[T GoAction] func(GoActionInternals) *T

type GoAction interface {
	Execute()
}

type GoActionInternals struct {
	*instance.ActionInstance
}

func NewGoActionInternals(actionName string) GoActionInternals {
	return GoActionInternals{
		ActionInstance: instance.NewActionInstance(actionName),
	}
}
