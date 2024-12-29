package action

type GoActionRegistration[Action GoAction, Props any] struct {
	Constructor GoActionConstructor[Action, Props]
  DefaultProps Props
}

type GoActionConstructor[T GoAction, Props any] func(*ActionInstance) *T

type GoAction interface {
	Execute()
}
