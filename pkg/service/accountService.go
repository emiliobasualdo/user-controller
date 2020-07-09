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

func EditAccount(accId uint, newAcc Account) (Account, error)  {
	acc, err := GetAccountById(accId)
	if err != nil {
		return Account{}, err
	}
	return persistence.EditAccount(acc, newAcc)
}

func GetEnabledInstrumentsByAccountId(id uint) ([]Instrument, error) {
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

func ExecuteTransaction(originId uint, trans Transaction) (Transaction, error) {
	// legal transaction
	if trans.Amount <= 0 {
		return Transaction{}, &IllegalTransactionError{Message: "The amount has to be bigger than 0"}
	}
	// origin exists
	originAcc, err := GetAccountById(originId)
	if err != nil {
		return Transaction{}, err
	}
	// origin is enabled
	if !originAcc.DisabledSince.IsZero() {
		return Transaction{}, &DisabledAccountError{Message: "The origin account is disabled"}
	}
	// origin has enough balance
	if originAcc.Balance < trans.Amount {
		return Transaction{}, &NotEnoughCreditError{}
	}
	// instrument exists & belongs & is enabled
	oInstruments, err := GetEnabledInstrumentsByAccountId(originAcc.ID)
	if err != nil {
		return Transaction{}, err
	}
	instrumentFound := false
	for _, ins := range oInstruments {
		if ins.ID == trans.InstrumentId {
			instrumentFound = true
			break
		}
	}
	if !instrumentFound {
		return Transaction{}, &NoSuchInstrumentError{ID: trans.InstrumentId}
	}
	// destination exists
	destAcc, err := GetAccountById(trans.DestinationAccountId)
	if err != nil {
		return Transaction{}, err
	}
	// destination is enabled
	if !destAcc.DisabledSince.IsZero() {
		return Transaction{}, &DisabledAccountError{Message: "The destination account is disabled"}
	}
	// execute transaction
	fullTransaction := Transaction{
		Amount:               trans.Amount,
		InstrumentId:         trans.InstrumentId,
		OriginAccountId:      originAcc.ID,
		DestinationAccountId: destAcc.ID,
	}
	return persistence.ExecuteTransaction(originAcc, destAcc, fullTransaction)
}

func GetTransactions(accId uint) ([]Transaction, error) {
	return persistence.GetTransactions(accId)
}

