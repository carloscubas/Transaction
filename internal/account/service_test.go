package account

import (
	_ "database/sql"
	"testing"
	"time"
	"transaction/internal/config"
	"transaction/test"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmrobles/h2go"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init (){
	sc, _ := config.LoadServiceConfig("../../configs/dev.yaml")
	database = sc.Db.Database
	connection = sc.Db.Connection
}

func TestInsertTransaction(t *testing.T) {

	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)

	transaction := Transaction{
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          26.0,
		EventData:       time.Now(),
	}

	service := NewService(repository)
	response, _ := service.InsertTransaction(transaction)

	if response.Amount != -26.0 {
		t.Errorf("expected %f, got %f", -26.0, response.Amount)
	}
	setup.After()
}

func TestInsertAccount(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)

	account := Account{
		DocumentNumber: "578945558",
	}

	service := NewService(repository)
	response, _ := service.InsertAccount(account)

	if response.DocumentNumber != "578945558" {
		t.Errorf("expected %s, got %s", "578945558", response.DocumentNumber)
	}

	setup.After()
}

func TestGetAccount(t *testing.T) {

	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)

	service := NewService(repository)
	response, _ := service.GetAccount(1)

	if response.DocumentNumber != "123456" {
		t.Errorf("expected %s, got %s", "123456", response.DocumentNumber)
	}

	setup.After()
}

func TestGetOperationTypes(t *testing.T) {

	setup := test.NewDbTestConfig(database, connection)
	setup.Before()

	repository, _ := NewRepository(setup.Conn)

	service := NewService(repository)
	response, _ := service.GetOperationsType()

	if  len(response) != 4{
		t.Errorf("expected %d, got %d", 4, len(response))
	}

	setup.After()
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
