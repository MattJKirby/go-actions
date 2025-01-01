package testHelpers

import "go-actions/ga/action"

type TestActionValid struct {}
func (ta TestActionValid) Execute(){}

type TestActionNoExecute struct{}

func GetEmptyConstructor[Test action.GoAction]() action.GoActionConstructor[Test] {
	return func(*action.ActionInstance) *Test {
		return new(Test)
	}
}


