package account

import "strings"

const (
	CREDT         = "CREDIT"
	DEBIT     = "DEBIT"
)

// Service struct to hold repository
type Service struct {
	config Config
	repo   Repository
}

// NewService create service struct
func NewService(config Config, repository Repository) *Service {
	return &Service{
		config: config,
		repo:   repository,
	}
}

func (s Service) insertTransaction(transaction Transaction) error {

	operation, err := s.repo.GetOperationType(transaction.OperationTypeId)
	if err != nil {
		return err
	}

	amount := checkTypeOperation(operation.Type, transaction.Amount)
	transaction.Amount = amount

	err = s.repo.InsertTransactions(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) insertAccount(account Account) error {
	err := s.repo.InsertAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) getAccount(id int64) (*Account, error) {
	account, err := s.repo.GetAccount(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func checkTypeOperation(typeOperator string, ammount float64) float64 {
	if strings.Compare(typeOperator, DEBIT) == 0 {
		ammount = -ammount
	}
	return ammount
}
