package account

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/jmrobles/h2go"
)

func TestInsertAccountRepository(t *testing.T) {
	conn := before()

	repository, _ := NewRepository(conn)

	account := Account{
		DocumentNumber: "123456789",
	}

	Account, err := repository.InsertAccount(account)
	if err != nil {
		log.Fatalf("Can't insert in H2 Database: %s", err)
	}

	if Account.Id != 1 {
		t.Errorf("expected %d, got %d", 1, Account.Id)
	}
	after(conn)
}

func TestInsertTransactionRepository(t *testing.T) {
	conn := before()
	repository, _ := NewRepository(conn)

	transaction := Transaction{
		Amount: 23.68,
		OperationTypeId: 1,
		AccountId: 1,
		EventData: time.Now(),
	}

	response, _ := repository.InsertTransactions(transaction)
	if response.Amount != 23.68 {
		t.Errorf("expected %f, got %f", -23.68, response.Amount)
	}
	after(conn)
}

func TestGetAccountRepository(t *testing.T) {
	conn := before()
	repository, _ := NewRepository(conn)

	response, _ := repository.GetAccount(1)
	if response.DocumentNumber != "123456" {
		t.Errorf("expected %s, got %s", "123456", response.DocumentNumber)
	}
	after(conn)
}

func TestGetOperationTypeRepository(t *testing.T) {

	conn := before()
	repository, _ := NewRepository(conn)
	response, _ := repository.GetOperationType(4)
	if response.Type == DEBIT {
		t.Errorf("expected %s, got %s", DEBIT, response.Type)
	}
	after(conn)
}

func before() *sql.DB{

	var ddl [11]string
	ddl[0] = "CREATE TABLE Accounts (Account_ID INTEGER auto_increment primary key, Document_Number varchar(100) NOT NULL);"
	ddl[1] = "INSERT INTO Accounts (Account_ID, Document_Number) VALUES (1, '123456');"
	ddl[2] = "CREATE TABLE OperationsTypes ( OperationsType_ID INTEGER auto_increment primary key, Description varchar(200) NOT NULL, OperationsType varchar(200) NOT NULL);"
	ddl[3] = "CREATE TABLE Transactions (Transaction_ID INTEGER auto_increment primary key, Account_ID INTEGER NOT NULL, OperationsType_ID INTEGER NOT NULL, Amount DOUBLE NOT NULL, EventDate DATE);"
	ddl[4] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_Account FOREIGN KEY (Account_ID) REFERENCES Accounts(Account_ID);"
	ddl[5] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_OperationsType FOREIGN KEY (OperationsType_ID) REFERENCES OperationsTypes(OperationsType_ID);"
	ddl[6] = "ALTER TABLE Accounts ADD CONSTRAINT Accounts_UN UNIQUE KEY (Document_Number);"
	ddl[7] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (1, 'COMPRA A VISTA','DEBIT');"
	ddl[8] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (2, 'COMPRA PARCELADA','DEBIT');"
	ddl[9] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (3, 'SAQUE','DEBIT');"
	ddl[10] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (4, 'PAGAMENTO','CREDIT');"

	conn, err := sql.Open("h2",  "h2://sa@localhost/test?mem=true&logging=debug")

	if err != nil {
		log.Fatalf("Can't connet to H2 Database: %s", err)
		panic(err)
	}

	for i := 0; i < len(ddl); i++ {
		response, err := conn.Exec(ddl[i])

		fmt.Println(response)
		if err != nil {
			log.Fatalf("Can't exec ddl commands: %s", err)
			panic(err)
		}
	}

	return conn
}

func after(conn *sql.DB){
	conn.Close()
}
