package application

import (
	"errors"
	"flag"
)

func InitializeApplication() error {
	if configuration.PostgresConfig != nil {
		return ConnectDatabase(GetConfig().PostgresConfig)
	} else {
		return errors.New("the postgres database has no configuration")
	}
}

func init() {
	flag.StringVar(&configPath, "c", "", "Specify the configuration file.")
}
