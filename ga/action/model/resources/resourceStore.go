package resources

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/utils/marshalling"
)

type marshalledResource[T any] struct {
	Id       string `json:"id"`
	Resource *T     `json:"attributes"`
}

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

func (rs *ResourceStore[T]) Get(name string) (*T, error) {
	if resource, exists := rs.resources[name]; exists {
		return resource, nil
	}
	return nil, fmt.Errorf("could not retrieve resource '%s'", name)
}

func (rs *ResourceStore[T]) GetOrDefault(name string, resourceFn func() *T) *T {
	if _, exists := rs.resources[name]; !exists {
		rs.resources[name] = resourceFn()
	}
	return rs.resources[name]
}

func (rs *ResourceStore[T]) MarshalJSON() ([]byte, error) {
	resourceList := make([]marshalledResource[T], 0, len(rs.resources))
	for name, res := range rs.resources {
		resourceList = append(resourceList, marshalledResource[T]{name, res})
	}

	return json.Marshal(resourceList)
}

func (rs *ResourceStore[T]) UnmarshalJSON(data []byte) error {
	rawResources := make([]marshalledResource[json.RawMessage], 0, len(rs.resources))
	if _, err := marshalling.StrictDecode(data, &rawResources); err != nil {
		return err
	}

	for _, raw := range rawResources {
		if _, exists := rs.resources[raw.Id]; !exists {
			return fmt.Errorf("error unmashalling: resource with identifier '%s' does not exist", raw.Id)
		}

		if _, err := marshalling.StrictDecode(*raw.Resource, rs.resources[raw.Id]); err != nil {
			return err
		}
	}
	return nil
}
