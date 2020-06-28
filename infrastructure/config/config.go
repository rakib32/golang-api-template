package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	App      AppConfig       `mapstructure:"app"`
	Database DatabaseConfig  `mapstructure:"database"`
}

// application specific config
type AppConfig struct {
	Port                int    `mapstructure:"port"`
}

// database specific config
type DatabaseConfig struct {
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	Name        string        `mapstructure:"name"`
	Username    string        `mapstructure:"username"`
	Password    string        `mapstructure:"password"`
	SslMode     string        `mapstructure:"ssl_mode"`
	MaxLifeTime time.Duration `mapstructure:"max_life_time"`
	MaxIdleConn int           `mapstructure:"max_idle_conn"`
	MaxOpenConn int           `mapstructure:"max_open_conn"`
	Debug       bool          `mapstructure:"debug"`
}

// c is the configuration instance
var c Config

// Get returns all configurations
func Get() Config {
	return c
}

// Load the config
func Load() error {
	viper.BindEnv("env")
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	consulURL := viper.GetString("consul_url")
	consulPath := viper.GetString("consul_path")
	if consulURL == "" || consulPath == "" {
		return errors.New("CONSUL_URL or CONSUL_PATH is missing from ENV")
	}

	// read config from remote consul
	viper.SetConfigType("yml")
	if err := viper.AddRemoteProvider("consul", consulURL, consulPath); err != nil {
		return err
	}
	if err := viper.ReadRemoteConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}
