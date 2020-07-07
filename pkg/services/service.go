package services

import (
	. "massimple.com/wallet-controller/pkg/models"
	"massimple.com/wallet-controller/pkg/persistence"
	"time"
)

func GetAccount(login Login) Account {
	acc := Account{
		PhoneNumber: login.PhoneNumber,
	}
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
	var resp []Instrument
	for _, inst := range insts {
		if !inst.DisabledAt.IsZero() {
			resp = append(resp, inst)
		}
	}
	return resp, nil
}

func InsertInstrumentById(instrumentDto InstrumentDto) (Instrument, error) {
	instrument := instrumentDto.Builder().Build()
	return persistence.CreateInstrument(instrumentDto.AccountID, instrument)
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
