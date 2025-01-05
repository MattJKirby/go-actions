package testHelpers

import "go-actions/ga/action"

type ActionValid struct{}
type ActionValidProps struct{}
func (ta ActionValid) Execute() {}

type ActionNoExecute struct{}

func GetEmptyConstructor[Test action.GoAction, TestProps action.GoActionProps]() action.GoActionConstructor[Test, TestProps] {
	return func(*action.ActionInstance, TestProps) *Test {
		return new(Test)
	}
}
