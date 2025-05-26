package common

import (
	"go-actions/ga/libs/uid"
)

type ModelProperty struct {
	Uid  uid.ResourceUid `json:"uid"`
	Name string          `json:"name"`
}

func NewModelProperty(modelUid uid.ResourceUid, propertyType string, name string) *ModelProperty {
	return &ModelProperty{
		Uid:  uid.NewUidBuilder().FromParent(modelUid).WithSubResource(propertyType).WithSubResourceId(name).Build(),
		Name: name,
	}
}

func (ap ModelProperty) GetId() string {
	return ap.Uid.FullUid()
}
