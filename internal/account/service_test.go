package account

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestInsertTransaction(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	timeNow := time.Now()

	transactionSaved := Transaction{
		AccountId: 1,
		OperationTypeId: 1,
		Amount: -26.0,
		EventData: timeNow,
	}

	transaction := Transaction{
		AccountId: 1,
		OperationTypeId: 1,
		Amount: 26.0,
		EventData: timeNow,
	}

	mockRepository := NewMockRepository(mockCtrl)
	mockRepository.EXPECT().InsertTransactions(transactionSaved).Return(&transactionSaved, nil).Times(1)
	mockRepository.EXPECT().GetOperationType(transaction.OperationTypeId).Return(&OperationType{Type: DEBIT}, nil).Times(1)

	service := NewService(mockRepository)
	response, _ := service.InsertTransaction(transaction)
	fmt.Println(response)
}

func TestCheckTypeOperation(t *testing.T) {
	transaction := Transaction{
		AccountId: 1,
		OperationTypeId: 1,
		Amount: 26.30,
	}

	result := checkTypeOperation(DEBIT, transaction)
	if result.Amount != -26.30 {
		t.Errorf("expected %f, got %f", -26.30, result.Amount)
	}
}
