package io

import (
	"go-actions/ga/libs/uid"
)

type ActionProperty struct {
	Uid       *uid.ResourceUid `json:"uid"`
	Name      string `json:"name"`
}

func NewActionProperty(modelUid *uid.ResourceUid, propertyType string, name string) *ActionProperty {
	return &ActionProperty{
		Uid:       uid.NewResourceUid(uid.WithParentUid(modelUid), uid.WithSubResource(propertyType), uid.WithSubResourceId(name)),
		Name:      name,
	}
}

func (ap ActionProperty) GetActionUid() string {
	return ap.Uid.GetBaseUid()
}

func (ap ActionProperty) GetResourceId() string {
	return ap.Uid.GetUid()
}
