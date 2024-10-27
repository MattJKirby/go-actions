package action

import "go-actions/ga/action/instance"

type GoActionConstructor[T GoAction] func(GoActionInternals) *T

type GoAction interface {
	Execute()
}

type GoActionInternals struct {
	Instance *instance.ActionInstance
}

func NewGoActionInternals(actionName string) GoActionInternals {
	return GoActionInternals{
		Instance: instance.NewActionInstance(actionName),
	}
}