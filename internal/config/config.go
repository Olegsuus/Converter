package config

import "github.com/spf13/viper"

type Config struct {
	Env    string `mapstructure:"env"`
	Server Server `mapstructure:"server"`
	Log    Log    `mapstructure:"logs"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type Log struct {
	LogFile string `mapstructure:"log_file"`
}

func MustLoad() (*Config, error) {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
