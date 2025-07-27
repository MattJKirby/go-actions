package action

import (
	"fmt"
)

type Action[T GoAction] struct {
	Definition T
	Instance   *ActionInstance
	typeDef    *TypeDefinition
}

func NewAction[T GoAction](typeDef *TypeDefinition, inst *ActionInstance) (*Action[T], error) {
	act, ok := typeDef.ActionValue.Interface().(T)
	if !ok {
		return nil, fmt.Errorf("new action does not match expected type")
	}

	// internals := NewInternals(inst)
	// act.InitInternals(internals)
	act.Init(inst)

	return &Action[T]{
		Instance:         inst,
		Definition:       act,
		typeDef:          typeDef,
	}, nil
}
