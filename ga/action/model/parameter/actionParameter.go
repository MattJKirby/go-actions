package parameter

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/libs/uid"
	"go-actions/ga/utils/marshalling"
)

type ActionParameter[T any] struct {
	name         string
	uid          uid.ResourceUid
	value        T
	defaultValue T
}

type marshalledActionParameter[T any] struct {
	Uid   string `json:"uid"`
	Name  string `json:"name"`
	Value T      `json:"value"`
}

func NewActionParameter[T any](modelUid uid.ResourceUid, name string, defaultValue T) *ActionParameter[T] {
	return &ActionParameter[T]{
		name:         name,
		uid:          uid.NewUidBuilder().FromParent(modelUid).WithSubResource("parameter").WithSubResourceId(name).Build(),
		value:        defaultValue,
		defaultValue: defaultValue,
	}
}

func (ap ActionParameter[T]) GetId() string {
	return ap.uid.FullUid()
}

func (ap *ActionParameter[T]) Value() T {
	return ap.value
}

func (ap *ActionParameter[T]) DefaultValue() T {
	return ap.defaultValue
}

func (ap *ActionParameter[T]) SetValue(value T) {
	ap.value = value
}

func (ap *ActionParameter[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&marshalledActionParameter[T]{
		Uid:   ap.uid.FullUid(),
		Name:  ap.name,
		Value: ap.value,
	})
}

func (ap *ActionParameter[T]) UnmarshalJSON(data []byte) error {
	var marshalled marshalledActionParameter[T]
	if _, err := marshalling.StrictDecode(data, &marshalled); err != nil {
		return err
	}

	if marshalled.Name != ap.name {
		return fmt.Errorf("failed to unmarshal action parameter: '%s': name '%s' does not match expected '%s'", ap.name, marshalled.Name, ap.name)
	}

	ap.SetValue(marshalled.Value)
	return nil
}
