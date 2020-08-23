package service

import (
	"fmt"
	guuid "github.com/google/uuid"
	"massimple.com/wallet-controller/internal/dtos"
	. "massimple.com/wallet-controller/internal/models"
	"massimple.com/wallet-controller/internal/persistence"
	"strings"
)

func GetAccount(phoneNumber PhoneNumber) (Account, error) {
	// we parse the phone number
	pnString := phoneNumber.String()
	pnString = strings.ReplaceAll(pnString,"+", "00" )
	pnString = strings.ReplaceAll(pnString,"-", "" )
	pnString = strings.ReplaceAll(pnString," ", "" )
	pnString = strings.ReplaceAll(pnString,"(", "" )
	pnString = strings.ReplaceAll(pnString,")", "" )
	// we generate an id if it is required by persistence
	newID := generateId(PhoneNumber(pnString))
	return persistence.GetAccountByPhoneNumberOrCreate(phoneNumber, newID)
}

func generateId(number PhoneNumber) ID {
	idString := fmt.Sprintf("%s-%s", guuid.New(), number.String())
	return ID(idString)
}

func GetAccountById(id ID) (Account, error) {
	return persistence.GetAccountById(id)
}

func EditAccount(accId ID, newAcc dtos.AccountDto) error {
	orignal, err := GetAccountById(accId)
	if err != nil {
		return err
	}
	return persistence.ReplaceAccount(dtos.FillAccountFromDto(orignal,newAcc))
}

func GetEnabledInstrumentsByAccountId(id ID) ([]Instrument, error) {
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

func InsertInstrumentById(accId ID, instrument Instrument) error {
	acc, err := GetAccountById(accId)
	if err != nil {
		return err
	}
	acc.Instruments = append(acc.Instruments, instrument)
	return persistence.ReplaceAccount(acc)
}

func ExecuteTransaction(originId ID, trans Transaction) error {
	// legal transaction
	if trans.Amount <= 0 {
		return &IllegalTransactionError{Message: "The amount has to be bigger than 0"}
	}
	// origin exists
	originAcc, err := GetAccountById(originId)
	if err != nil {
		return err
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
		return &NoSuchInstrumentError{ID: trans.InstrumentId.String()}
	}
	// destination exists
	destAcc, err := GetAccountById(trans.DestinationAccountId)
	if err != nil {
		return err
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

func GetTransactions(accId ID) ([]Transaction, error) {
	acc, err := GetAccountById(accId)
	if err != nil {
		return []Transaction{}, err
	}
	return acc.Transactions, nil
}

