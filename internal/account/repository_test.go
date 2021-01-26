package account

import (
	"log"
	"testing"
	"time"
	"transaction/internal/config"
	"transaction/test"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmrobles/h2go"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var database string
var connection string

func init (){
	sc, _ := config.LoadServiceConfig("../../configs/dev.yaml")
	database = sc.Db.Database
	connection = sc.Db.Connection
}

func TestInsertAccountRepository(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)

	setup.Before()

	repository, _ := NewRepository(setup.Conn)

	account := Account{
		DocumentNumber: "123456789",
	}

	Account, err := repository.InsertAccount(account)
	if err != nil {
		log.Fatalf("Can't insert in Database: %s", err)
	}

	if Account.Id != 1 {
		t.Errorf("expected %d, got %d", 1, Account.Id)
	}

	setup.After()
}

func TestInsertTransactionRepository(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)
	transaction := Transaction{
		Amount:          23.68,
		OperationTypeId: 1,
		AccountId:       1,
		EventData:       time.Now(),
	}

	response, _ := repository.InsertTransaction(transaction)
	if response.Amount != 23.68 {
		t.Errorf("expected %f, got %f", -23.68, response.Amount)
	}
	setup.After()

}

func TestGetAccountRepository(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)
	response, _ := repository.GetAccount(1)
	if response.DocumentNumber != "123456" {
		t.Errorf("expected %s, got %s", "123456", response.DocumentNumber)
	}

	setup.After()
}

func TestGetOperationTypeRepository(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)
	response, _ := repository.GetOperationType(4)
	if response.Type == DEBIT {
		t.Errorf("expected %s, got %s", DEBIT, response.Type)
	}

	setup.After()
}

func TestGetOperationTypeListRepository(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)
	response, _ := repository.GetOperationTypes()

	if  len(response) != 4{
		t.Errorf("expected %d, got %d", 4, len(response))
	}

	setup.After()
}