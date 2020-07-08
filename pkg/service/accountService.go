package service

import (
	. "massimple.com/wallet-controller/pkg/models"
	"massimple.com/wallet-controller/pkg/persistence"
	"time"
)

func GetAccount(phoneNumber string) (Account, error) {
	acc := Account{PhoneNumber: phoneNumber}
	return persistence.GetAccountByPhoneNumberOrCreate(acc)
}

func GetAccountById(id uint) (Account, error) {
	return persistence.GetAccountById(id)
}

func GetInstrumentsById(id uint) ([]Instrument, error) {
	insts, err := persistence.GetInstrumentsByAccountId(id)
	if err != nil {
		return nil, err
	}
	// we remove disabled instruments
	resp := make([]Instrument, 0)
	for _, inst := range insts {
		if inst.DisabledAt.IsZero() {
			resp = append(resp, inst)
		}
	}
	return resp, nil
}

func InsertInstrumentById(accId uint, instrument Instrument) (Instrument, error) {
	return persistence.CreateInstrument(accId, instrument)
}

func DeleteInstrumentById(accountId uint, instrumentId uint) error {
	inst, err := persistence.GetInstrumentsById(instrumentId)
	if err != nil {
		return err
	}
	if inst.AccountID != accountId {
		return &UnauthorizedError{Message: "The instrument does not belong to the user"}
	}
	if inst.DisabledAt.IsZero() {
		inst.DisabledAt = time.Now()
		err = persistence.SaveInstrument(inst)
	}
	return err
}
