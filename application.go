package main

import (
	"github.com/apsdehal/go-logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"massimple.com/wallet-controller/internal/persistence"
	"massimple.com/wallet-controller/internal/service"
	"massimple.com/wallet-controller/internal/webapp"
	"os"
)


func main() {
	envInit()
	persistence.Init()
	service.Init()
	webapp.Serve()
}

func envInit() {
	log, _ := logger.New("Environment", 1, os.Stdout)
	// set defaults
	viper.SetConfigName("config-default")
	viper.AddConfigPath("./")		// search locally in this directory
	viper.AddConfigPath(".")		// search locally in this directory
	viper.AddConfigPath("$HOME")  	// when deployed search in this directory
	err := viper.ReadInConfig() 		// Find and read the config file
	if err != nil {
		panic(err)
	}
	// specific based on env
	value, exists := os.LookupEnv("ENV")
	if !exists {
		panic("Environment variable ENV must be set to PRO or DEV")
	}
	if value == "PROD" {
		log.Info("Running in PROD mode")
	} else {
		log.Info("Running in DEV mode")
		viper.SetConfigName("config-dev")
	}
	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}
}



