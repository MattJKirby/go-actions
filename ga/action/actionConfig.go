package action

import "github.com/google/uuid"

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (ac *Config) GenerateUid() string {
	return uuid.New().String()
}
