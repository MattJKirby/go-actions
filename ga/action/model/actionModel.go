package model

import (
	"fmt"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/store"
)

type ActionModel struct {
	ActionName string                          `json:"name"`
	ActionUid  string                          `json:"uid"`
	Parameters *store.PropertyStore[any]       `json:"parameters"`
	Inputs     *store.PropertyStore[io.Input]  `json:"inputs"`
	Outputs    *store.PropertyStore[io.Output] `json:"outputs"`
}

type ActionConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: store.NewPropertyStore[any](),
		Inputs:     store.NewPropertyStore[io.Input](),
		Outputs:    store.NewPropertyStore[io.Output](),
	}
}
