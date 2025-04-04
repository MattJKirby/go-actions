package config

import (
	"go-actions/ga/utils/packageConfig"
	"go-actions/ga/utils/uid"
)

type GlobalConfig struct {
	UidGenerator uid.UidGenerator
}

func DefaultGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		UidGenerator: &uid.DefaultUidGenerator{},
	}
}

func WithCustomUidGenerator(gen uid.UidGenerator) packageConfig.Option[GlobalConfig] {
	return func(gc *GlobalConfig) {
		gc.UidGenerator = gen
	}
}
