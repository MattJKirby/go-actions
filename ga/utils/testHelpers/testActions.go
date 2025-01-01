package testHelpers

import "go-actions/ga/action"

type ActionValid struct {}
func (ta ActionValid) Execute(){}

type ActionNoExecute struct{}

func GetEmptyConstructor[Test action.GoAction]() action.GoActionConstructor[Test] {
	return func(*action.ActionInstance) *Test {
		return new(Test)
	}
}


