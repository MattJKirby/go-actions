package config

import (
	"go-actions/ga/libs/packageConfig"
	"go-actions/ga/utils/uid"
)

type GlobalConfig struct {
	UidProvider uid.Provider
}

func DefaultGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		UidProvider: &uid.DefaultProvider{},
	}
}

func WithCustomUidProvider(gen uid.Provider) packageConfig.Option[GlobalConfig] {
	return func(gc *GlobalConfig) {
		gc.UidProvider = gen
	}
}
