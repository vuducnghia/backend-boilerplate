package application

import (
	"backend-boilerplate/logger"
	"encoding/json"
	"errors"
	"flag"
	"github.com/rs/zerolog"

	"os"
	"strings"
)

var (
	configPath    string
	configuration *Configuration

	ConfigFileError       = errors.New("error reading the configuration file")
	ConfigFileFormatError = errors.New("configuration file appears to be in the wrong format")
)

type Configuration struct {
	ApplicationConfig *SystemConfig    `json:"application_config"`
	PostgresConfig    *DatabaseConfig  `json:"database_config"`
	DirectoryConfig   *DirectoryConfig `json:"directory_config"`
}

type SystemConfig struct {
	Secret   string `json:"secret"`
	LogLevel string `json:"log_level"`
	IsDebug  bool   `json:"is_debug"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	LogLevel string `json:"log_level"`
}

type DirectoryConfig struct {
	BaseAssetUrl         string `json:"base_asset_url"`
	BaseUploadsDirectory string `json:"base_uploads_directory"`
	ImagesDirectory      string `json:"images_directory"`
}

func setDefaultConfig() *Configuration {
	return &Configuration{
		ApplicationConfig: &SystemConfig{
			LogLevel: "info",
			Secret:   "CHANGEMENOW",
		},
		DirectoryConfig: &DirectoryConfig{
			BaseAssetUrl:         "localhost",
			BaseUploadsDirectory: "uploads/",
			ImagesDirectory:      "images/",
		},
		PostgresConfig: nil,
	}
}

func GetConfig() *Configuration {
	if configuration == nil {
		log.Warn().Msg("WARNING: configuration was not initialized")
		return &Configuration{}
	}
	return configuration
}

func LoadConfig() error {
	flag.Parse()
	configuration = setDefaultConfig()
	if strings.HasSuffix(configPath, ".json") {
		configFile, err := os.Open(configPath)
		if err != nil {
			return ConfigFileError
		}
		defer func(f *os.File) {
			cErr := f.Close()
			if cErr != nil {
				log.Error().Err(cErr).Msg("error occurred closing the json config file")
			}
		}(configFile)

		decoder := json.NewDecoder(configFile)
		err = decoder.Decode(configuration)
		if err != nil {
			return ConfigFileFormatError
		}
	}

	lvl, err := zerolog.ParseLevel(configuration.ApplicationConfig.LogLevel)

	if err != nil {
		configuration.ApplicationConfig.LogLevel = "info"
		lvl = zerolog.InfoLevel
		log.Error().Err(err).Msg("log level in the application configuration was incorrect")
	}
	log.SetGlobalLogLevel(lvl.String())
	return nil
}
