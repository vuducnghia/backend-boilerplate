package application

import (
	"errors"
	"flag"
)

func InitializeApplication() error {
	if configuration.PostgresConfig != nil {
		ConnectDatabase(GetConfig().PostgresConfig)
	} else if configuration.RedisConfig != nil {
		NewRedisCache(GetConfig().RedisConfig)
	} else {
		return errors.New("the postgres database has no configuration")
	}
	return nil
}

func init() {
	flag.StringVar(&configPath, "c", "", "Specify the configuration file.")
}
