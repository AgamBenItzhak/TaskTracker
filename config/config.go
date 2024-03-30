package config

import (
	"github.com/AgamBenItzhak/TaskTracker/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var ServerConfig *Config
var CfgFileFlag string
var Flags pflag.FlagSet

// InitConfig reads environment variables and config file if set
func init() {
	ServerConfig = NewConfig()
	ServerConfig.ApplyConfig()
	Flags = *pflag.NewFlagSet("config", pflag.ExitOnError)
}

func NewConfig() *Config {
	return &Config{}
}

// ApplyConfig reads and applies all configurations in the following order of priority:
// 1. Command line flags
// 2. Environment variables
// 3. Config file
// 4. Default values
// If a config file is specified, it will be read from the specified path
// Otherwise, the config file will be read from the default path
func (c *Config) ApplyConfig() {
	SetDefaultConfig()

	if CfgFileFlag != "" {
		viper.SetConfigFile(CfgFileFlag)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.TaskTracker")
		viper.AddConfigPath("/etc/TaskTracker/")
	}

	if err := viper.ReadInConfig(); err != nil {
		logger.LogService.Error("Failed to read config file: ", err)
	} else {
		logger.LogService.Info("Reading configurations from file: ", viper.ConfigFileUsed())
	}

	viper.SetEnvPrefix("TASKTRACKER")
	viper.AutomaticEnv()
	BindEnv()

	// Add flags to viper to override config file values if set
	viper.BindPFlags(&Flags)

	if err := viper.Unmarshal(c); err != nil {
		logger.LogService.Error("Failed to load configurations: ", err)
		return
	}

	logger.LogService.Info("Configurations loaded successfully")
}

func SetDefaultConfig() {
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.dbname", "tasktracker")
	viper.SetDefault("db.sslmode", "disable")
	viper.SetDefault("log.level", "info")
	viper.SetDefault("jwt.secret", "secret")
	viper.SetDefault("jwt.expiration", "3600")
}

func BindEnv() {
	viper.BindEnv("server.host", "TASKTRACKER_SERVER_HOST")
	viper.BindEnv("server.port", "TASKTRACKER_SERVER_PORT")
	viper.BindEnv("db.host", "TASKTRACKER_DB_HOST")
	viper.BindEnv("db.port", "TASKTRACKER_DB_PORT")
	viper.BindEnv("db.user", "TASKTRACKER_DB_USER")
	viper.BindEnv("db.password", "TASKTRACKER_DB_PASSWORD")
	viper.BindEnv("db.dbname", "TASKTRACKER_DB_DBNAME")
	viper.BindEnv("db.sslmode", "TASKTRACKER_DB_SSLMODE")
	viper.BindEnv("log.level", "TASKTRACKER_LOG_LEVEL")
	viper.BindEnv("jwt.secret", "TASKTRACKER_JWT_SECRET")
	viper.BindEnv("jwt.expiration", "TASKTRACKER_JWT_EXPIRATION")
}
