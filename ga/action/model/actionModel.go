package model

import (
	"fmt"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/resources"
)

type ActionModel struct {
	ActionName string                              `json:"name"`
	ActionUid  string                              `json:"uid"`
	Parameters *resources.ResourceStore[any]       `json:"parameters"`
	Inputs     *resources.ResourceStore[io.Input]  `json:"inputs"`
	Outputs    *resources.ResourceStore[io.Output] `json:"outputs"`
}

type ActionModelConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionModelConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: resources.NewResourceStore[any](),
		Inputs:     resources.NewResourceStore[io.Input](),
		Outputs:    resources.NewResourceStore[io.Output](),
	}
}
