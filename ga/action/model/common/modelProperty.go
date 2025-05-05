package common

import (
	"go-actions/ga/libs/uid"
)

type ModelProperty struct {
	Uid  *uid.ResourceUid `json:"uid"`
	Name string           `json:"name"`
}

func NewModelProperty(modelUid *uid.ResourceUid, propertyType string, name string) *ModelProperty {
	return &ModelProperty{
		Uid:  uid.NewResourceUid(uid.WithParentUid(modelUid), uid.WithSubResource(propertyType), uid.WithSubResourceId(name)),
		Name: name,
	}
}

func (ap ModelProperty) GetActionUid() string {
	return ap.Uid.GetBaseUid()
}

func (ap ModelProperty) GetResourceId() string {
	return ap.Uid.GetUid()
}
