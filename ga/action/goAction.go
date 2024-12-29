package action

type GoActionRegistration[Action GoAction, Props any] struct {
	Constructor GoActionConstructor[Action]
  Props Props
}

type GoActionConstructor[T GoAction] func(*ActionInstance) *T

type GoAction interface {
	Execute()
}
