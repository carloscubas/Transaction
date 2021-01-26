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

func TestRepository(t *testing.T) {
	sc, _ := config.LoadServiceConfig("../../configs/dev.yaml")
	setup := test.NewDbTestConfig(sc.Db.Database, sc.Db.Connection)
	repository, _ := NewRepository(setup.Conn)
	setup.Before()

	t.Run("TestInsertAccountRepository", func(t *testing.T) {

		account := Account{
			DocumentNumber: "123456789",
		}

		Account, err := repository.InsertAccount(account)
		if err != nil {
			log.Fatalf("Can't insert in Database: %s", err)
		}

		if Account.DocumentNumber != "123456789" {
			t.Errorf("expected %s, got %s", "123456789", Account.DocumentNumber)
		}

	})
	t.Run("TestInsertTransactionRepository", func(t *testing.T) {
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
	})
	t.Run("TestGetAccountRepository", func(t *testing.T) {
		response, _ := repository.GetAccount(1)
		if response.DocumentNumber != "123456" {
			t.Errorf("expected %s, got %s", "123456", response.DocumentNumber)
		}
	})
	t.Run("TestGetOperationTypeRepository", func(t *testing.T) {
		response, _ := repository.GetOperationType(4)
		if response.Type == DEBIT {
			t.Errorf("expected %s, got %s", DEBIT, response.Type)
		}

	})
	t.Run("TestGetOperationTypeListRepository", func(t *testing.T) {
		response, _ := repository.GetOperationTypes()
		if  len(response) != 4{
			t.Errorf("expected %d, got %d", 4, len(response))
		}
	})

	setup.After()
	setup.Conn.Close()
}