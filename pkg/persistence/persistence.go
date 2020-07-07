package persistence

import (
	"github.com/apsdehal/go-logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "massimple.com/wallet-controller/pkg/models"
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
	db.Table("accounts").CreateTable(&Account{})
	db.Table("instruments").CreateTable(&Instrument{})
}

func GetAccountByPhoneNumberOrCreate(query Account) Account {
	account, err := getByPhoneNumber(query)
	if err == gorm.ErrRecordNotFound {
		account = query
		db.Create(&account)
		log.InfoF("Creating user with phone number: %s", account.PhoneNumber)
	}
	return account
}

func getByPhoneNumber(query Account) (Account, error) {
	var account Account
	if db.Where(&Account{PhoneNumber: query.PhoneNumber}).First(&account).RecordNotFound() {
		return account, gorm.ErrRecordNotFound
	}
	return account, nil
}

func GetInstrumentsByAccountId(id uint) ([]Instrument, error)  {
	acc, err := GetAccountById(id)
	if err != nil {
		return nil, err
	}
	var insts []Instrument
	if err := db.Model(&acc).Related(&insts).Error; err != nil {
		return []Instrument{}, err
	}
	return insts, nil
}

func GetInstrumentsById(id uint) (Instrument, error) {
	var inst Instrument
	err := db.First(&inst, id).Error
	return inst, err
}
func CreateInstrument(accountId uint, inst Instrument) (Instrument, error) {
	acc, err := GetAccountById(accountId)
	if err != nil {
		return Instrument{}, err
	}
	if err := db.Model(&acc).Association("Instruments").Append(&inst).Error; err != nil {
		return Instrument{}, err
	}
	return inst, err
}

func SaveInstrument(instrument Instrument) error {
	if err := db.Save(&instrument).Error; err != nil {
		return err
	}
	return nil
}

func GetAccountById(id uint) (Account, error) {
	var acc Account
	if db.First(&acc, id).RecordNotFound() {
		return Account{}, &NoSuchAccountError{ID: id}
	}
	return acc, nil
}