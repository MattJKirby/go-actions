package parameter

import (
	"go-actions/ga/action/model/resources"
)

type Store struct {
	*resources.ResourceStore[any]
}

func NewStore() *Store {
	return &Store{
		resources.NewResourceStore[any](),
	}
}
