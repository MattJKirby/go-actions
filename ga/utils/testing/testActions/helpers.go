package testActions

import "go-actions/ga/action"

func GenerateRegistration[T action.GoAction](act T) action.GoActionRegistration[T] {
	return action.GoActionRegistration[T]{
		Action:       act,
	}
}
