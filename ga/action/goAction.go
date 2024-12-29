package action

type GoActionRegistration[Type GoAction, Props any] struct {
	Props Props
	Constructor GoActionConstructor[Type]
}

type GoActionConstructor[T GoAction] func(*ActionInstance) *T

type GoAction interface {
	Execute()
}
