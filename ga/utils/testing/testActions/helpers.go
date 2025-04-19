package testActions

import "go-actions/ga/action"

func GenerateRegistration[T action.GoAction, P action.GoActionProps](act T, props P) action.GoActionRegistration[T, P] {
	return action.GoActionRegistration[T, P]{
		Action:       act,
		DefaultProps: props,
	}
}
