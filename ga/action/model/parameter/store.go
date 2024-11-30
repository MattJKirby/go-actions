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

func GetOrDefault[T any](name string, defaultValue T) func(*Store) *ActionParameter[T] {
	return func(s *Store) *ActionParameter[T] {
		defaultAsAny := any(NewActionParameter(name, defaultValue))
		got := s.GetOrDefault(name, &defaultAsAny)
		return (*got).(*ActionParameter[T])
	}
}
