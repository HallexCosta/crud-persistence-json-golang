package config

// AppConfig ...
type AppConfig struct {
	Persistence Persistence
}

var _AppConfig = new(AppConfig)

// ImportAppConfig ...
func ImportAppConfig() *AppConfig {
	return _AppConfig
}
