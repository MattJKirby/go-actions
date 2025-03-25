package store

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/utils/marshalling"
)

type marshalledProperty[PropertyType any] struct {
	PropertyId string        `json:"propertyId"`
	Property   *PropertyType `json:"property"`
}

type PropertyStore[PropertyType any] struct {
	properties map[string]*PropertyType
}

func NewPropertyStore[PropertyType any]() *PropertyStore[PropertyType] {
	return &PropertyStore[PropertyType]{
		properties: make(map[string]*PropertyType),
	}
}

func (rs *PropertyStore[T]) Add(name string, property *T) {
	rs.properties[name] = property
}

func (rs *PropertyStore[T]) Get(name string) (*T, error) {
	if property, exists := rs.properties[name]; exists {
		return property, nil
	}
	return nil, fmt.Errorf("could not retrieve property '%s'", name)
}

func (rs *PropertyStore[T]) GetOrDefault(name string, propertyFn func() *T) *T {
	if _, exists := rs.properties[name]; !exists {
		rs.properties[name] = propertyFn()
	}
	return rs.properties[name]
}

func (rs *PropertyStore[T]) MarshalJSON() ([]byte, error) {
	properties := make([]marshalledProperty[T], 0, len(rs.properties))
	for name, res := range rs.properties {
		properties = append(properties, marshalledProperty[T]{name, res})
	}

	return json.Marshal(properties)
}

func (rs *PropertyStore[T]) UnmarshalJSON(data []byte) error {
	rawProperty := make([]marshalledProperty[json.RawMessage], 0, len(rs.properties))
	if _, err := marshalling.StrictDecode(data, &rawProperty); err != nil {
		return err
	}

	for _, raw := range rawProperty {
		if _, exists := rs.properties[raw.PropertyId]; !exists {
			return fmt.Errorf("error unmashalling: property with identifier '%s' does not exist", raw.PropertyId)
		}

		if _, err := marshalling.StrictDecode(*raw.Property, rs.properties[raw.PropertyId]); err != nil {
			return err
		}
	}
	return nil
}
