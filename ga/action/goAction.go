package action

import "go-actions/ga/action/instance"

type GoActionConstructor[T GoAction] func() *T

type GoAction interface {
	Execute()
}

type GoActionInternals struct {
	instance *instance.ActionInstance
}

func NewGoActionInternals(actionName string) *GoActionInternals {
	return &GoActionInternals{
		instance: instance.NewActionInstance(actionName),
	}
}