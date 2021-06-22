package config

import "github.com/spf13/viper"

func Load() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	SetConfigDefaults()

	return nil
}

func SetConfigDefaults() {
	viper.SetDefault("server.port", "1337")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.user", "boilerplate")
	viper.SetDefault("database.password", "boilerplate")
	viper.SetDefault("database.name", "boilerplate")
	viper.SetDefault("database.port", "5432")
}
