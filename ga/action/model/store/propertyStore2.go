package store

// import (
// 	"encoding/json"
// 	"fmt"
// 	"go-actions/ga/utils/marshalling"
// )

// // type marshalledProperty[PropertyType any] struct {
// // 	PropertyId string        `json:"propertyId"`
// // 	Property   *PropertyType `json:"property"`
// // }

// type IdentifiableProperty interface {
// 	GetPropertyId() string
// }

// type PropertyStore[PropertyType IdentifiableProperty] struct {
// 	properties map[string]*PropertyType
// }

// func NewPropertyStore[PropertyType IdentifiableProperty]() *PropertyStore[PropertyType] {
// 	return &PropertyStore[PropertyType]{
// 		properties: make(map[string]*PropertyType),
// 	}
// }

// func (rs *PropertyStore[T]) Add(property *T) {
// 	rs.properties[(*property).GetPropertyId()] = property
// }

// func (rs *PropertyStore[T]) Get(name string) (*T, error) {
// 	if property, exists := rs.properties[name]; exists {
// 		return property, nil
// 	}
// 	return nil, fmt.Errorf("could not retrieve property '%s'", name)
// }

// func (rs *PropertyStore[T]) GetOrDefault(name string, propertyFn func() *T) *T {
// 	if _, exists := rs.properties[name]; !exists {
// 		rs.properties[name] = propertyFn()
// 	}
// 	return rs.properties[name]
// }

// func (rs *PropertyStore[T]) MarshalJSON() ([]byte, error) {
// 	properties := make([]*T, 0, len(rs.properties))
// 	for _, res := range rs.properties {
// 		properties = append(properties, res)
// 	}
// 	return json.Marshal(properties)
// }

// func (rs *PropertyStore[T]) UnmarshalJSON(data []byte) error {
// 	rawProperty := make([]json.RawMessage, 0, len(rs.properties))
// 	if _, err := marshalling.StrictDecode(data, &rawProperty); err != nil {
// 		return err
// 	}

// 	for _, raw := range rawProperty {
// 		property := *new(T)
// 		if _, err := marshalling.StrictDecode(raw, &property); err != nil {
// 			return err
// 		}

// 		if _, exists := rs.properties[property.GetPropertyId()]; !exists {
// 			return fmt.Errorf("error unmashalling: property with identifier '%s' does not exist", property.GetPropertyId())
// 		}

// 		rs.properties[property.GetPropertyId()] = &property
// 	}
// 	return nil
// }
