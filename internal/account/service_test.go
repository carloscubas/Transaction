package account

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmrobles/h2go"
)

func TestInsertTransaction(t *testing.T) {
	conn := Before()
	defer conn.Close()

	repository, _ := NewRepository(conn)

	transaction := Transaction{
		AccountId: 1,
		OperationTypeId: 1,
		Amount: 26.0,
		EventData: time.Now(),
	}

	service := NewService(repository)
	response, _ := service.InsertTransaction(transaction)

	if response.Amount != -26.0 {
		t.Errorf("expected %f, got %f", -26.0, response.Amount)
	}

}

func TestInsertAccount(t *testing.T) {
	conn := Before()
	defer conn.Close()

	repository, _ := NewRepository(conn)

	account := Account{
		DocumentNumber: "578945558",
	}

	service := NewService(repository)
	response, _ := service.InsertAccount(account)

	if response.DocumentNumber != "578945558" {
		t.Errorf("expected %s, got %s", "578945558", response.DocumentNumber)
	}
}

func TestGetAccount(t *testing.T) {
	conn := Before()
	defer conn.Close()

	repository, _ := NewRepository(conn)

	service := NewService(repository)
	response, _ := service.GetAccount(1)

	if response.DocumentNumber != "123456" {
		t.Errorf("expected %s, got %s", "123456", response.DocumentNumber)
	}
}


func TestCheckTypeOperation(t *testing.T) {
	transaction := Transaction{
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          26.30,
	}

	result := checkTypeOperation(DEBIT, transaction)
	if result.Amount != -26.30 {
		t.Errorf("expected %f, got %f", -26.30, result.Amount)
	}
}
