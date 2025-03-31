package io

import (
	"fmt"
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

type Config interface {
	GenerateUid() string
}

func NewActionReference(config Config, sourceUid string, targetUid string) *ActionReference {
	return &ActionReference{
		ReferenceUid:    fmt.Sprintf("ref:%s", config.GenerateUid()),
		SourceActionUid: sourceUid,
		TargetActionUid: targetUid,
	}
}

func (ar ActionReference) GetPropertyId() string {
	return ar.ReferenceUid
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
