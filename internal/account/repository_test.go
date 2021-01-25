package account

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"
	"transaction/internal/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmrobles/h2go"
)

func TestInsertAccountRepository(t *testing.T) {
	conn := Before()

	defer conn.Close()

	repository, _ := NewRepository(conn)

	account := Account{
		DocumentNumber: "123456789",
	}

	Account, err := repository.InsertAccount(account)
	if err != nil {
		log.Fatalf("Can't insert in Database: %s", err)
	}

	if Account.Id != 2 {
		t.Errorf("expected %d, got %d", 1, Account.Id)
	}

}

func TestInsertTransactionRepository(t *testing.T) {
	conn := Before()
	defer conn.Close()

	repository, _ := NewRepository(conn)

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
}

func TestGetAccountRepository(t *testing.T) {
	conn := Before()
	defer conn.Close()

	repository, _ := NewRepository(conn)

	response, _ := repository.GetAccount(1)
	if response.DocumentNumber != "123456" {
		t.Errorf("expected %s, got %s", "123456", response.DocumentNumber)
	}

}

func TestGetOperationTypeRepository(t *testing.T) {

	conn := Before()
	defer conn.Close()

	repository, _ := NewRepository(conn)
	response, _ := repository.GetOperationType(4)
	if response.Type == DEBIT {
		t.Errorf("expected %s, got %s", DEBIT, response.Type)
	}

}
func TestGetOperationTypeListRepository(t *testing.T) {

	conn := Before()
	defer conn.Close()

	repository, _ := NewRepository(conn)
	response, _ := repository.GetOperationTypes()

	if  len(response) != 4{
		t.Errorf("expected %d, got %d", 4, len(response))
	}

}


func Before() *sql.DB {
	var ddl [17]string
	ddl[0] = "DROP TABLE IF EXISTS Transactions;"
	ddl[1] = "DROP TABLE IF EXISTS OperationsTypes;"
	ddl[2] = "DROP TABLE IF EXISTS Accounts;"
	ddl[3] = "CREATE TABLE Accounts (Account_ID INTEGER auto_increment primary key, Document_Number varchar(100) NOT NULL);"
	ddl[4] = "CREATE TABLE OperationsTypes ( OperationsType_ID INTEGER auto_increment primary key, Description varchar(200) NOT NULL, OperationsType varchar(200) NOT NULL);"
	ddl[5] = "CREATE TABLE Transactions (Transaction_ID INTEGER auto_increment primary key, Account_ID INTEGER NOT NULL, OperationsType_ID INTEGER NOT NULL, Amount DOUBLE NOT NULL, EventDate DATE);"
	ddl[6] = "Delete from Transactions;"
	ddl[7] = "Delete from OperationsTypes;"
	ddl[8] = "Delete from Accounts;"
	ddl[9] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_Account FOREIGN KEY (Account_ID) REFERENCES Accounts(Account_ID);"
	ddl[10] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_OperationsType FOREIGN KEY (OperationsType_ID) REFERENCES OperationsTypes(OperationsType_ID);"
	ddl[11] = "ALTER TABLE Accounts ADD CONSTRAINT Accounts_UN UNIQUE KEY (Document_Number);"
	ddl[12] = "INSERT INTO Accounts (Document_Number) VALUES ('123456');"
	ddl[13] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (1, 'COMPRA A VISTA','DEBIT');"
	ddl[14] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (2, 'COMPRA PARCELADA','DEBIT');"
	ddl[15] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (3, 'SAQUE','DEBIT');"
	ddl[16] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (4, 'PAGAMENTO','CREDIT');"

	var database string
	var dbConnection string

	if len(os.Getenv("API_DB_DATABASE")) > 0 {
		database = os.Getenv("API_DB_DATABASE")
		dbConnection = os.Getenv("API_DB_CONNECTION")
	} else {
		sc, _ := config.LoadServiceConfig("../../configs/dev.yaml")
		database = sc.Db.Database
		dbConnection = sc.Db.Connection
	}

	conn, err := sql.Open(database, dbConnection)

	if err != nil {
		log.Fatalf("Can't connet to H2 Database: %s", err)
	}

	tx, _ := conn.Begin()
	for i := 0; i < len(ddl); i++ {
		tx.Exec(ddl[i])
	}
	tx.Commit()

	return conn
}
