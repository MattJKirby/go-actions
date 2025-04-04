package testActions

import "go-actions/ga/action"

func GenerateEmptyCtor[T action.GoAction, P action.GoActionProps]() action.GoActionConstructor[T, P] {
	return func(*action.ActionInstance, P) *T {
		return new(T)
	}
}

func GenerateRegistration[T action.GoAction, P action.GoActionProps](ctor action.GoActionConstructor[T, P], props *P) action.GoActionRegistration[T, P] {
	return action.GoActionRegistration[T, P]{
		Constructor:  ctor,
		DefaultProps: props,
	}
}
