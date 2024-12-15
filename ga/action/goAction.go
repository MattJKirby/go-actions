package action

type GoActionConstructor[T GoAction] func(*ActionInstance) *T

type GoAction interface {
	Execute()
}
