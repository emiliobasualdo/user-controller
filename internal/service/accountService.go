package service

import (
	"massimple.com/wallet-controller/internal/dtos"
	. "massimple.com/wallet-controller/internal/models"
	"massimple.com/wallet-controller/internal/persistence"
)

func GetAccount(phoneNumber string) (Account, error) {
	acc := Account{PhoneNumber: phoneNumber}
	return persistence.GetAccountByPhoneNumberOrCreate(acc)
}

func GetAccountById(id string) (Account, error) {
	return persistence.GetAccountById(id)
}

func EditAccount(accId string, newAcc dtos.AccountDto) error {
	orignal, err := GetAccountById(accId)
	if err != nil {
		return err
	}
	return persistence.ReplaceAccount(dtos.FillAccountFromDto(orignal,newAcc))
}

func GetEnabledInstrumentsByAccountId(id string) ([]Instrument, error) {
	acc, err := persistence.GetAccountById(id)
	if err != nil {
		return nil, err
	}
	// we remove disabled instruments
	insts := acc.Instruments
	resp := make([]Instrument, 0)
	for _, inst := range insts {
		if !inst.Disabled {
			resp = append(resp, inst)
		}
	}
	return resp, nil
}

func InsertInstrumentById(accId string, instrument Instrument) error {
	acc, err := GetAccountById(accId)
	if err != nil {
		return err
	}
	acc.Instruments = append(acc.Instruments, instrument)
	return persistence.ReplaceAccount(acc)
}

func ExecuteTransaction(originId string, trans Transaction) error {
	// legal transaction
	if trans.Amount <= 0 {
		return &IllegalTransactionError{Message: "The amount has to be bigger than 0"}
	}
	// origin exists
	originAcc, err := GetAccountById(originId)
	if err != nil {
		return err
	}
	// origin is enabled
	if !originAcc.Disabled {
		return &DisabledAccountError{Message: "The origin account is disabled"}
	}
	// origin has enough balance
	if originAcc.Balance < trans.Amount {
		return &NotEnoughCreditError{}
	}
	// instrument exists & belongs & is enabled
	oInstruments, err := GetEnabledInstrumentsByAccountId(originAcc.ID)
	if err != nil {
		return err
	}
	instrumentFound := false
	for _, ins := range oInstruments {
		if ins.ID == trans.InstrumentId {
			instrumentFound = true
			break
		}
	}
	if !instrumentFound {
		return &NoSuchInstrumentError{ID: trans.InstrumentId}
	}
	// destination exists
	destAcc, err := GetAccountById(trans.DestinationAccountId)
	if err != nil {
		return err
	}
	// destination is enabled
	if !destAcc.Disabled {
		return &DisabledAccountError{Message: "The destination account is disabled"}
	}
	// execute transaction
	fullTransaction := Transaction{
		Amount:               trans.Amount,
		InstrumentId:         trans.InstrumentId,
		OriginAccountId:      originAcc.ID,
		DestinationAccountId: destAcc.ID,
	}
	return persistence.SaveTransaction(fullTransaction)
}

func GetTransactions(accId string) ([]Transaction, error) {
	acc, err := GetAccountById(accId)
	if err != nil {
		return []Transaction{}, err
	}
	return acc.Transactions, nil
}

