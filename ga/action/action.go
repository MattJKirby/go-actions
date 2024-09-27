package action

type ActionConstructor[T any] func () T

type ActionFunction interface {
	Execute()
}

type Action struct {
	def *ActionDefinition
}

func NewAction(def *ActionDefinition) *Action {

	return &Action{
		def,
	}
}
