package config

import (
	"strings"

	"github.com/spf13/viper"
)

// ReadConfig reads the config in from the game server
func ReadConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {

	if defaults == nil {
		defaults = map[string]interface{}{
			"server": map[string]interface{}{
				"port":     8080,
				"hostname": "localhost",
				"list":     false,
			},
			"logging": map[string]interface{}{
				"debug":  false,
				"format": "json",
			},
		}
	}

	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetEnvPrefix("GOSERVER")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.WatchConfig()
	err := v.ReadInConfig()
	return v, err
}
