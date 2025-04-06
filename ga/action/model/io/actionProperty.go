package io

import "fmt"

type ActionProperty struct {
	actionUid string
	Uid  string
	Name string
}

func NewActionProperty(actionUid string, propertyType string, name string) *ActionProperty {
	return &ActionProperty{
		actionUid: actionUid,
		Uid:  fmt.Sprintf("%s:%s:%s", actionUid, propertyType, name),
		Name: name,
	}
}

func (ap ActionProperty) GetActionUid() string {
	return ap.actionUid
}

func (ap ActionProperty) GetPropertyId() string {
	return ap.Uid
}
