package io

import (
	"fmt"
	"go-actions/ga/app/config"
)

type ReferencableProperty interface {
	GetActionUid() string
	GetPropertyId() string
}

type ReferencableSource interface {
	ReferencableProperty
	AssignTargetReference(*PartialActionReference) error
}

type ReferencableTarget interface {
	ReferencableProperty
	AssignSourceReference(*PartialActionReference) error
}

type ActionReference struct {
	referenceUid string
	source       ReferencableSource
	target       ReferencableTarget
}

type PartialActionReference struct {
	ReferenceUid string `json:"ReferenceUid"`
	ActionUid    string `json:"ActionUid"`
}

func NewActionReference(globalConfig *config.GlobalConfig, source ReferencableSource, target ReferencableTarget) *ActionReference {
	return &ActionReference{
		referenceUid: fmt.Sprintf("ref:%s", globalConfig.UidGenerator.GenerateUid()),
		source:       source,
		target:       target,
	}
}

func (par PartialActionReference) GetPropertyId() string {
	return par.ReferenceUid
}

func (ar *ActionReference) getSourceReference() *PartialActionReference {
	return &PartialActionReference{
		ReferenceUid: ar.referenceUid,
		ActionUid:    ar.source.GetActionUid(),
	}
}

func (ar *ActionReference) getTargetReference() *PartialActionReference {
	return &PartialActionReference{
		ReferenceUid: ar.referenceUid,
		ActionUid:    ar.target.GetActionUid(),
	}
}

func (ar *ActionReference) AssignReferences() error {
	if err := ar.source.AssignTargetReference(ar.getTargetReference()); err != nil {
		return err
	}
	if err := ar.target.AssignSourceReference(ar.getSourceReference()); err != nil {
		return err
	}
	return nil
}
