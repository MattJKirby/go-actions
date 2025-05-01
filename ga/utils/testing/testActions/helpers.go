package testActions

import "go-actions/ga/action"

func GenerateRegistration[T action.GoAction](act T) action.ActionRegistration[T] {
	return action.ActionRegistration[T]{
		Action: act,
	}
}
