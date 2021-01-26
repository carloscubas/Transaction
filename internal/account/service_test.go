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

func TestService(t *testing.T) {
	sc, _ := config.LoadServiceConfig("../../configs/dev.yaml")
	setup := test.NewDbTestConfig(sc.Db.Database, sc.Db.Connection)
	repository, _ := NewRepository(setup.Conn)
	setup.Before()

	t.Run("TestInsertTransaction", func(t *testing.T) {

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

	})
	t.Run("TestInsertAccount", func(t *testing.T) {

		account := Account{
			DocumentNumber: "578945558",
		}

		service := NewService(repository)
		response, _ := service.InsertAccount(account)

		if response.DocumentNumber != "578945558" {
			t.Errorf("expected %s, got %s", "578945558", response.DocumentNumber)
		}

	})
	t.Run("TestGetAccount", func(t *testing.T) {

		service := NewService(repository)
		response, _ := service.GetAccount(1)

		if response.DocumentNumber != "123456" {
			t.Errorf("expected %s, got %s", "123456", response.DocumentNumber)
		}

	})
	t.Run("TestGetOperationTypes", func(t *testing.T) {

		service := NewService(repository)
		response, _ := service.GetOperationsType()

		if  len(response) != 4{
			t.Errorf("expected %d, got %d", 4, len(response))
		}

	})
	t.Run("TestCheckTypeOperation", func(t *testing.T) {
		transaction := Transaction{
			AccountId:       1,
			OperationTypeId: 1,
			Amount:          26.30,
		}

		result := checkTypeOperation(DEBIT, transaction)
		if result.Amount != -26.30 {
			t.Errorf("expected %f, got %f", -26.30, result.Amount)
		}
	})

	setup.After()
	setup.Conn.Close()
}