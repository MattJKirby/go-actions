package model

import (
	"fmt"
	"go-actions/ga/action/model/io"
)

type ActionModel struct {
	ActionName string                              `json:"name"`
	ActionUid  string                              `json:"uid"`
	Parameters *PropertyStore[any]       `json:"parameters"`
	Inputs     *PropertyStore[io.Input]  `json:"inputs"`
	Outputs    *PropertyStore[io.Output] `json:"outputs"`
}

type ActionModelConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionModelConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: NewPropertyStore[any](),
		Inputs:     NewPropertyStore[io.Input](),
		Outputs:    NewPropertyStore[io.Output](),
	}
}
