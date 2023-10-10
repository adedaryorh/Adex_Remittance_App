package utils

import "github.com/spf13/viper"

type Config struct {
	DBdriver string `mapstructure:"DB_DRIVER""`
	DB_source string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (c *Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadConfig()

	if err
}
