package resources

import (
	"encoding/json"
	"fmt"
)

type ResourceStore[T any] struct {
	resources map[string]*T
}

func NewResourceStore[T any]() *ResourceStore[T] {
	return &ResourceStore[T]{
		resources: make(map[string]*T),
	}
}

func (rs *ResourceStore[T]) Add(name string, resource *T) {
	rs.resources[name] = resource
}

func (rs *ResourceStore[T]) GetOrDefault(name string, resource *T) *T {
	_, exists := rs.resources[name]
	if !exists {
		rs.resources[name] = resource
	}
	return rs.resources[name]
}

func (rs *ResourceStore[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(rs.resources)
}

func (rs *ResourceStore[T]) UnmarshalJSON(data []byte) error {
	rawResources := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &rawResources); err != nil {
		return err
	}

	for name, raw := range rawResources {
		if _, exists := rs.resources[name]; !exists {
			return fmt.Errorf("error unmashalling: resource '%s' does not exist", name)
		}

		if err := json.Unmarshal(raw, rs.resources[name]); err != nil {
			return err
		}
	}
	return nil
}
