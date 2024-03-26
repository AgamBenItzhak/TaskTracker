package config

import (
	"fmt"
	"os"

	"github.com/AgamBenItzhak/TaskTracker/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CfgFile string

func initLogger() {
	err := logger.InitLogger()
	if err != nil {
		fmt.Println("Failed to initialize logger due to error: ", err)
		os.Exit(1)
	}
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	initLogger()
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".TaskTracker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".TaskTracker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.LogService.Info("Using config file: ", viper.ConfigFileUsed())
	}
}
