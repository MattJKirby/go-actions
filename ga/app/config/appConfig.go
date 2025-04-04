package config

import "go-actions/ga/utils/packageConfig"


type ApplicationConfig struct {
	Global *GlobalConfig
}

func DefaultApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{
		Global: packageConfig.NewPackageConfig(DefaultGlobalConfig()),
	}
}

func NewAppConfig(opts ...packageConfig.Option[ApplicationConfig]) *ApplicationConfig {
	return packageConfig.NewPackageConfig(DefaultApplicationConfig(), opts...)
}

func WithGlobalConfig(globalOpts ...packageConfig.Option[GlobalConfig]) packageConfig.Option[ApplicationConfig] {
	return func(ac *ApplicationConfig) {
		ac.Global = packageConfig.NewPackageConfig(DefaultGlobalConfig(), globalOpts...)
	}
}
