package io

import "fmt"

type ActionProperty struct {
	Uid string
	Name string
}

func NewActionProperty(actionUid string, propertyType string, name string) *ActionProperty {
	return &ActionProperty{
		Uid: fmt.Sprintf("%s:%s:%s", actionUid, propertyType, name),
		Name: name,
	}
}

func (ap ActionProperty) GetPropertyId() string {
	return ap.Uid
}

