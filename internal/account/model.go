package account

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Repository interface {
	InsertAccount(account Account) (*Account, error)
	UpdateAvailableCreditLimitAccount(account *Account) error
	InsertTransaction(transaction Transaction) (*Transaction, error)
	GetAccount(id int64) (*Account, error)
	GetOperationType(id int64) (*OperationType, error)
	GetOperationTypes() ([]OperationType, error)
	GetTransactions(id int64) ([]Transaction, error)
	SetBalance(transaction Transaction) error
}

// Account is a structure that represents the Account request.
type Account struct {
	Id             int64  `json:"Account_ID"`
	DocumentNumber string `json:"DocumentNumber"`
	AvailableCreditLimit float64   `json:"AvailableCreditLimit"`
}

func (t Account) validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.DocumentNumber, validation.Required),
	)
}

// OperationType is a structure that represents the OperationType request.
type OperationType struct {
	Id          int64  `json:"OperationType_ID"`
	Description string `json:"Description"`
	Type        string `json:"OperationType"`
}

// Transaction is a structure that represents the Transaction request.
type Transaction struct {
	Id              int64     `json:"TransactionID"`
	AccountId       int64     `json:"AccountID"`
	OperationTypeId int64     `json:"OperationsTypeID"`
	Amount          float64   `json:"Amount"`
	Balance          float64   `json:"Balance"`
	EventData       time.Time `json:"EventDate"`
}

func (t Transaction) validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Amount, validation.Required, validation.Min(0.0)),
		validation.Field(&t.AccountId, validation.Required),
		validation.Field(&t.OperationTypeId, validation.Required),
	)
}

type TransactionResponse struct {
	DocumentNumber string  `json:"Document"`
	Description    string  `json:"Description"`
	Amount         float64 `json:"Amount"`
}
