package account

import "strings"

const (
	DEBIT = "DEBIT"
)

// Service struct to hold repository
type Service struct {
	repo Repository
}

// NewService create service struct
func NewService(repository Repository) *Service {
	return &Service{
		repo: repository,
	}
}

// InsertTransaction insert the new transaction
func (s Service) InsertTransaction(transaction Transaction) (*Transaction, error) {

	operation, err := s.repo.GetOperationType(transaction.OperationTypeId)
	if err != nil {
		return nil, err
	}

	transaction = checkTypeOperation(operation.Type, transaction)

	tra, err := s.repo.InsertTransaction(transaction)
	if err != nil {
		return nil, err
	}
	return tra, nil
}

// InsertAccount insert the new account
func (s Service) InsertAccount(account Account) (*Account, error) {
	ac, err := s.repo.InsertAccount(account)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

// GetAccount get the account
func (s Service) GetAccount(id int64) (*Account, error) {
	account, err := s.repo.GetAccount(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// GetOperationsType get the operation type
func (s Service) GetOperationsType() ([]OperationType, error) {
	types, err := s.repo.GetOperationTypes()
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (s Service) GetTransactions() ([]TransactionResponse, error) {
	transactions, err := s.repo.GetTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil

}

// checkTypeOperation check is amount is positive or negative
func checkTypeOperation(typeOperator string, transaction Transaction) Transaction {
	if strings.Compare(typeOperator, DEBIT) == 0 {
		transaction.Amount = -transaction.Amount
	}
	return transaction
}
