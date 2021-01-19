package account

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Repository interface {
	InsertAccount(account Account) error
	InsertTransactions(transaction Transaction) error
	GetAccount(id int64) (*Account, error)
	GetOperationType(id int64) (*OperationType, error)
}

type Account struct {
	Id             int64  `json:"Account_ID"`
	DocumentNumber string `json:"Document_Number"`
}

func (t Account) validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.DocumentNumber, validation.Required),
	)
}

type OperationType struct {
	Id          int64  `json:"OperationType_ID"`
	Description string `json:"Description"`
	Type        string `json:"OperationType"`
}

type Transaction struct {
	Id              int64     `json:"TransactionID"`
	AccountId       int64     `json:"AccountID"`
	OperationTypeId int64     `json:"OperationTypeID"`
	Amount          float64   `json:"Amount"`
	EventData       time.Time `json:"EventDate"`
}

func (t Transaction) validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Amount, validation.Required, validation.Min(0.0)),
		validation.Field(&t.AccountId, validation.Required),
		validation.Field(&t.OperationTypeId, validation.Required),
	)
}
