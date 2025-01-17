package model

import "github.com/google/uuid"

type ModelConfig struct{}

func NewModelConfig() *ModelConfig {
	return &ModelConfig{}
}

func (ic *ModelConfig) GenerateUid() string {
	return uuid.New().String()
}
