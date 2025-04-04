package io

import (
	"fmt"
	"go-actions/ga/app/config"
)

type ActionReference struct {
	ReferenceUid    string `json:"ReferenceUid"`
	SourceActionUid string `json:"SourceActionUid"`
	TargetActionUid string `json:"TargetActionUid"`
}

type PartialActionReference struct {
	ReferenceUid string `json:"ReferenceUid"`
	ActionUid    string `json:"ActionUid"`
}

func NewActionReference(globalConfig *config.GlobalConfig, sourceUid string, targetUid string) *ActionReference {
	return &ActionReference{
		ReferenceUid:    fmt.Sprintf("ref:%s", globalConfig.UidGenerator.GenerateUid()),
		SourceActionUid: sourceUid,
		TargetActionUid: targetUid,
	}
}

func (par PartialActionReference) GetPropertyId() string {
	return par.ReferenceUid
}

func (ar *ActionReference) GetSourceReference() *PartialActionReference {
	return &PartialActionReference{
		ReferenceUid: ar.ReferenceUid,
		ActionUid:    ar.SourceActionUid,
	}
}

func (ar *ActionReference) GetTargetReference() *PartialActionReference {
	return &PartialActionReference{
		ReferenceUid: ar.ReferenceUid,
		ActionUid:    ar.TargetActionUid,
	}
}
