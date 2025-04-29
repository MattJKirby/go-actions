package config

import "go-actions/ga/libs/packageConfig"

type ApplicationConfig struct {
	Global *GlobalConfig
}

func DefaultApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{
		Global: packageConfig.NewPackageConfig(DefaultGlobalConfig()),
	}
}

func WithGlobalConfigOptions(globalOpts ...packageConfig.Option[GlobalConfig]) packageConfig.Option[ApplicationConfig] {
	return func(ac *ApplicationConfig) {
		ac.Global = packageConfig.NewPackageConfig(ac.Global, globalOpts...)
	}
}
