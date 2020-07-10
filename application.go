package main

import (
	"github.com/apsdehal/go-logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"massimple.com/wallet-controller/pkg/persistence"
	"massimple.com/wallet-controller/pkg/service"
	"massimple.com/wallet-controller/pkg/webapp"
	"os"
)

var log *logger.Logger

func main() {
	var err error
	log, err = logger.New("Persistence", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	envInit()
	persistence.Init()
	service.Init()
	webapp.Serve(viper.GetString("server.port"))
}

func envInit() {
	// set defaults
	viper.SetConfigName("config-default")
	viper.AddConfigPath("./")		// search locally in this directory
	viper.AddConfigPath("$HOME")  	// when deployed search in this directory
	err := viper.ReadInConfig() 		// Find and read the config file
	if err != nil {
		panic(err)
	}
	// specific based on env
	value, exists := os.LookupEnv("ENV")
	if exists && value == "PROD" {
		log.Info("Running in PROD mode")
		viper.SetConfigName("config-prod")
	} else {
		log.Info("Running in DEV mode")
		viper.SetConfigName("config-dev")
	}
	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}
}


