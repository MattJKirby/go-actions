package config

import "go-actions/ga/utils/uid"

type Option func(*GlobalConfig)

type GlobalConfig struct {
	UidGenerator uid.UidGenerator
}

func NewGlobalConfig(opts ...Option) *GlobalConfig {
	cfg := &GlobalConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithCustomUidGenerator(gen uid.UidGenerator) Option {
	return func(gc *GlobalConfig) {
		gc.UidGenerator = gen
	}
}

