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

func (s Service) InsertTransaction(transaction Transaction) (*Transaction, error) {

	operation, err := s.repo.GetOperationType(transaction.OperationTypeId)
	if err != nil {
		return nil, err
	}

	transaction = checkTypeOperation(operation.Type, transaction)

	tra, err := s.repo.InsertTransactions(transaction)
	if err != nil {
		return nil, err
	}
	return tra, nil
}

func (s Service) InsertAccount(account Account) (*Account, error) {
	ac, err := s.repo.InsertAccount(account)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (s Service) GetAccount(id int64) (*Account, error) {
	account, err := s.repo.GetAccount(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s Service) GetOperationsType() ([]OperationType, error) {
	types, err := s.repo.GetOperationTypes()
	if err != nil {
		return nil, err
	}
	return types, nil
}

func checkTypeOperation(typeOperator string, transaction Transaction) Transaction {
	if strings.Compare(typeOperator, DEBIT) == 0 {
		transaction.Amount = -transaction.Amount
	}
	return transaction
}
