package config

import "go-actions/ga/utils/config"


type ApplicationConfig struct {
	Global *GlobalConfig
}

func DefaultApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{
		Global: config.NewPackageConfig(DefaultGlobalConfig()),
	}
}

func NewAppConfig(opts ...config.Option[ApplicationConfig]) *ApplicationConfig {
	return config.NewPackageConfig(DefaultApplicationConfig(), opts...)
}

func WithGlobalConfig(globalOpts ...config.Option[GlobalConfig]) config.Option[ApplicationConfig] {
	return func(ac *ApplicationConfig) {
		ac.Global = config.NewPackageConfig(DefaultGlobalConfig(), globalOpts...)
	}
}
