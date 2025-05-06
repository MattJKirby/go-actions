package common

import (
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
)

type ActionReference struct {
	Uid uid.ResourceUid `json:"uid"`
	Source uid.ResourceUid `json:"source"`
	Target uid.ResourceUid `json:"target"`
}

type ResourceReference struct {
	Uid uid.ResourceUid `json:"uid"`
	Resource uid.ResourceUid `json:"resource"`
}

func NewActionReference(globalConfig *config.GlobalConfig, source uid.ResourceUid, target uid.ResourceUid) *ActionReference {
	return &ActionReference{
		Uid: uid.NewUidBuilder().WithResource("Ref").WithUid(globalConfig.UidGenerator.GenerateUid()).Build(),
		Source: source,
		Target: target,
	}
}

func (ar *ActionReference) GetSourceReference() *ResourceReference {
	return &ResourceReference{
		Uid: ar.Uid,
		Resource: ar.Source,
	}
}

func (ar *ActionReference) GetTargetReference() *ResourceReference {
	return &ResourceReference{
		Uid: ar.Uid,
		Resource: ar.Target,
	}
}

func (ar ActionReference) GetResourceId() string {
	return ar.Uid.FullUid()
}


func (rr ResourceReference) GetResourceId() string {
	return rr.Uid.FullUid()
}