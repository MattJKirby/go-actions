package app

import (
	"go-actions/ga/app/config"
	"go-actions/ga/libs/packageConfig"
)

type ApplicationConfig struct {
	Global *config.GlobalConfig
}

func DefaultApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{
		Global: packageConfig.NewPackageConfig(config.DefaultGlobalConfig()),
	}
}

func WithGlobalConfigOptions(globalOpts ...packageConfig.Option[config.GlobalConfig]) packageConfig.Option[ApplicationConfig] {
	return func(ac *ApplicationConfig) {
		ac.Global = packageConfig.NewPackageConfig(ac.Global, globalOpts...)
	}
}
