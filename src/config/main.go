package config

// AppConfig ...
type AppConfig struct{}

var _AppConfig = new(AppConfig)

// NewAppConfig ...
func NewAppConfig() *AppConfig {
	return _AppConfig
}
