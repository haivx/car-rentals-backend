package config

import "github.com/spf13/viper"

type Config struct {
	HOST          string `mapstructure:"HOST"`
	PORT          string `mapstructure:"PORT"`
	USERNAME      string `mapstructure:"USERNAME"`
	PASSWORD      string `mapstructure:"PASSWORD"`
	DATABASE_NAME string `mapstructure:"DATABASE_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
