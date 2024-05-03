package application

import (
	"backend-boilerplate/logger"
	"encoding/json"
	"errors"
	"flag"
	"github.com/rs/zerolog"
	"os"
	"strconv"
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
	RedisConfig       *RedisConfig     `json:"redis_config"`
	DirectoryConfig   *DirectoryConfig `json:"directory_config"`
}

type SystemConfig struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	RequestTimeout int    `json:"request_timeout"`
	LogLevel       string `json:"log_level"`
	IsDebug        bool   `json:"is_debug"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	IsDebug  bool   `json:"is_debug"`
}

type DirectoryConfig struct {
	BaseAssetUrl         string `json:"base_asset_url"`
	BaseUploadsDirectory string `json:"base_uploads_directory"`
	ImagesDirectory      string `json:"images_directory"`
}

func setDefaultConfig() *Configuration {
	return &Configuration{
		ApplicationConfig: &SystemConfig{
			LogLevel:     "info",
			AccessToken:  "CHANGEMENOW",
			RefreshToken: "CHANGEMENOW",
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
	loadEnvironment(configuration)
	lvl, err := zerolog.ParseLevel(configuration.ApplicationConfig.LogLevel)
	if err != nil {
		configuration.ApplicationConfig.LogLevel = "info"
		lvl = zerolog.InfoLevel
		log.Error().Err(err).Msg("log level in the application configuration was incorrect")
	}
	log.SetGlobalLogLevel(lvl.String())
	return nil
}

func checkEnvironment(key string, original string) string {
	if val, ok := os.LookupEnv(key); !ok {
		return original
	} else {
		return val
	}
}

func loadEnvironment(c *Configuration) {
	if c.ApplicationConfig == nil {
		c.ApplicationConfig = &SystemConfig{}
	}
	c.ApplicationConfig.AccessToken = checkEnvironment("application_config__access_token", c.ApplicationConfig.AccessToken)
	c.ApplicationConfig.RefreshToken = checkEnvironment("application_config__refresh_token", c.ApplicationConfig.RefreshToken)
	c.ApplicationConfig.LogLevel = checkEnvironment("application_config__log_level", c.ApplicationConfig.LogLevel)
	if val, ok := os.LookupEnv("application_config__is_debug"); ok {
		if bVal, pErr := strconv.ParseBool(val); pErr == nil {
			c.ApplicationConfig.IsDebug = bVal
		}
	}

	if c.PostgresConfig == nil {
		c.PostgresConfig = &DatabaseConfig{}
	}
	c.PostgresConfig.Port = checkEnvironment("database_config__port", c.PostgresConfig.Port)
	c.PostgresConfig.Host = checkEnvironment("database_config__host", c.PostgresConfig.Host)
	c.PostgresConfig.Username = checkEnvironment("database_config__username", c.PostgresConfig.Username)
	c.PostgresConfig.Password = checkEnvironment("database_config__password", c.PostgresConfig.Password)
	c.PostgresConfig.Database = checkEnvironment("database_config__database", c.PostgresConfig.Database)
	if val, ok := os.LookupEnv("database_config__is_debug"); ok {
		if bVal, pErr := strconv.ParseBool(val); pErr == nil {
			c.PostgresConfig.IsDebug = bVal
		}
	}

	if c.DirectoryConfig == nil {
		c.DirectoryConfig = &DirectoryConfig{}
	}
	c.DirectoryConfig.BaseAssetUrl = checkEnvironment("directory_config__base_asset_url", c.DirectoryConfig.BaseAssetUrl)
	c.DirectoryConfig.BaseUploadsDirectory = checkEnvironment("directory_config__base_uploads_directory", c.DirectoryConfig.BaseUploadsDirectory)
	c.DirectoryConfig.ImagesDirectory = checkEnvironment("directory_config__images_directory", c.DirectoryConfig.ImagesDirectory)
}
