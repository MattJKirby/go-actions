package io

import (
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
)

type ResourceReference struct {
	Uid      uid.ResourceUid  `json:"uid"`
	Source   *uid.ResourceUid `json:"source,omitempty"`
	Target   *uid.ResourceUid `json:"target,omitempty"`
	Resource *uid.ResourceUid `json:"resource,omitempty"`
}

func NewActionReference(globalConfig *config.GlobalConfig, source *uid.ResourceUid, target *uid.ResourceUid) *ResourceReference {
	return &ResourceReference{
		Uid:    uid.NewUidBuilder().WithResource("Ref").WithUid(globalConfig.UidProvider.New()).Build(),
		Source: source,
		Target: target,
	}
}

func (ar *ResourceReference) GetSourceReference() *ResourceReference {
	return &ResourceReference{
		Uid:    ar.Uid,
		Source: ar.Source,
	}
}

func (ar *ResourceReference) GetTargetReference() *ResourceReference {
	return &ResourceReference{
		Uid:    ar.Uid,
		Target: ar.Target,
	}
}

func (rr ResourceReference) GetId() string {
	return rr.Uid.FullUid()
}
