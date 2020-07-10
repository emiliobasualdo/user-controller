package persistence

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "massimple.com/wallet-controller/pkg/models"
)

var tables = []struct {
	TableName  	string
	Model 		interface{}
}{
	{"accounts", &Account{}},
	{"instruments", &Instrument{}},
	{"transactions", &Transaction{}},
}

func GetAccountByPhoneNumberOrCreate(query Account) (Account, error) {
	account, err := getByPhoneNumber(query)
	if err == nil {
		return account, nil
	}
	if err != gorm.ErrRecordNotFound {
		return Account{}, err
	}
	log.InfoF("Creating user with phone number: %s", query.PhoneNumber)
	account = query
	if err := db.Create(&account).Error; err == nil {
		return account, nil
	} else {
		return Account{}, nil
	}

}

func GetAccountById(id uint) (Account, error) {
	var acc Account
	if db.Model(&acc).Preload("Instruments").First(&acc, id).RecordNotFound() {
		return Account{}, &NoSuchAccountError{ID: id}
	}
	return acc, nil
}

func getByPhoneNumber(query Account) (Account, error) {
	var account Account
	if db.Where(&Account{PhoneNumber: query.PhoneNumber}).First(&account).RecordNotFound() {
		return account, gorm.ErrRecordNotFound
	}
	return account, nil
}

func EditAccount(old Account, new Account) (Account, error) {
	err := db.Model(&old).Update(new).Error
	return old, err
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

func ExecuteTransaction(originAcc Account, destAcc Account, trans Transaction)  (Transaction, error){
	// transaction start
	tx := db.Begin()
	// origin subtract balance
	oErr := tx.Model(&originAcc).Update(Account{
		Balance: originAcc.Balance - trans.Amount,
	}).Error
	// destination add balance
	dErr := tx.Model(&destAcc).Update(Account{
		Balance: destAcc.Balance + trans.Amount,
	}).Error
	// create new transaction
	tErr := tx.Create(&trans).Error
	// if any error -> rollback
	if oErr != nil || dErr != nil || tErr != nil {
		tx.Rollback()
		return Transaction{}, pickFirstNonNil(oErr, dErr, tErr)
	}
	// commit
	tx.Commit()
	return trans, nil
}

func GetTransactions(accId uint) ([]Transaction, error){
	var trans []Transaction
	err := db.Order("created_at desc").Find(&trans).Where(Transaction{OriginAccountId: accId}).Error
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func pickFirstNonNil(values ...error) error {
	for _, val := range values {
		if val != nil {
			return val
		}
	}
	return errors.New("unknown error")
}