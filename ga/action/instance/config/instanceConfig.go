package config

import "github.com/google/uuid"

type InstanceConfig struct{}

func NewInstanceConfig() *InstanceConfig {
	return &InstanceConfig{}
}

func (ic *InstanceConfig) GenerateUid() string {
	return uuid.New().String()
}
