package action

type GoActionConstructor[T GoAction] func() *T

type GoAction interface {
	Execute()
}