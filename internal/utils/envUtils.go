package utils

import (
	"github.com/apsdehal/go-logger"
	"github.com/spf13/viper"
	"os"
)


func EnvInit(env string) {
	_ = os.Setenv("ENV", env)
	log, _ := logger.New("Environment", 1, os.Stdout)
	// set defaults
	viper.SetConfigName("config-default")
	viper.AddConfigPath("./../../")		// search locally in this directory
	viper.AddConfigPath("./../../../")	// search locally in this directory
	viper.AddConfigPath(".")		// search locally in this directory
	viper.AddConfigPath("$HOME")  	// when deployed search in this directory
	err := viper.ReadInConfig() 		// Find and read the config file
	if err != nil {
		panic(err)
	}
	// specific based on env
	if env == "PROD" {
		log.Info("Running in PROD mode")
	} else if env == "DEV"{
		log.Info("Running in DEV mode")
		viper.SetConfigName("config-dev")
	} else {
		log.Info("Running in TEST mode")
		viper.SetConfigName("config-test")
	}
	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}
}