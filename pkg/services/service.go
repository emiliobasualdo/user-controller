package services

import (
	. "massimple.com/wallet-controller/pkg/models"
	"massimple.com/wallet-controller/pkg/persistence"
	"time"
)

func GetAccount(phoneNumber string) Account {
	acc := Account{PhoneNumber: phoneNumber}
	return persistence.GetAccountByPhoneNumberOrCreate(acc)
}

func getAccountById(id uint) (Account, error) {
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

func DeleteInstrumentById(instrumentId uint) error {
	inst, err := persistence.GetInstrumentsById(instrumentId)
	if err != nil {
		return err
	}
	if inst.DisabledAt.IsZero() {
		inst.DisabledAt = time.Now()
		err = persistence.SaveInstrument(inst)
	}
	return err
}
