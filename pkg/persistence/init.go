package persistence

import (
	"github.com/apsdehal/go-logger"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var log *logger.Logger
func Init(_log *logger.Logger) {
	log = _log
	_db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	log.Info("Db connected")
	db = _db
	db.LogMode(true)
	for _, table := range tables {
		db.Table(table.TableName).CreateTable(table.Model)
		db.AutoMigrate(table.Model)
	}
}
