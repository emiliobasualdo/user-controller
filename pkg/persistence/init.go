package persistence

import (
	"fmt"
	"github.com/apsdehal/go-logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)


var db *gorm.DB
var log *logger.Logger
func Init() {
	var err error
	log, err = logger.New("Persistence", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	// we get the basicUri
	basicUri := viper.GetString("database.connectionUri")
	// if aws provides env variables, we want to use those for prod uri
	dbName, exists := os.LookupEnv("RDS_DB_NAME")
	if exists {
		port 		:= os.Getenv("RDS_PORT")
		host 		:= os.Getenv("RDS_HOSTNAME")
		user 		:= os.Getenv("RDS_USERNAME")
		password	:= os.Getenv("RDS_PASSWORD")
		// user:password@host:port/dbname dbname must be previously created(CREATE DATABASE dbname)
		basicUri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	}
	// gorm needs some settings to interact with mysql
	conUri := fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", basicUri)
	// we now connect
	log.InfoF("Connecting to: %s", conUri)
	viper.Set("database.connectionUri", conUri)
	db, err = gorm.Open("mysql", conUri)
	if err != nil {
		panic(err)
	}
	log.Info("DB connected successfully")
	db.LogMode(viper.GetBool("database.verbose")) // todo env
	// we create/update the tables if not exist
	for _, table := range tables {
		db.Table(table.TableName).CreateTable(table.Model)
		db.AutoMigrate(table.Model)
	}
}
