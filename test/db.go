package test

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DbTestConfig struct {
	Conn *sql.DB
	ddlBefore [14]string
	ddlAfter [3]string
	dataBase string
	connection string
}

func NewDbTestConfig(dataBase string, connection string) *DbTestConfig {

	var ddlBefore [14]string
	ddlBefore[0] = "DROP TABLE IF EXISTS Transactions;"
	ddlBefore[1] = "DROP TABLE IF EXISTS OperationsTypes;"
	ddlBefore[2] = "DROP TABLE IF EXISTS Accounts;"
	ddlBefore[3] = "CREATE TABLE Accounts (Account_ID INTEGER auto_increment primary key, Document_Number varchar(100) NOT NULL);"
	ddlBefore[4] = "CREATE TABLE OperationsTypes ( OperationsType_ID INTEGER auto_increment primary key, Description varchar(200) NOT NULL, OperationsType varchar(200) NOT NULL);"
	ddlBefore[5] = "CREATE TABLE Transactions (Transaction_ID INTEGER auto_increment primary key, Account_ID INTEGER NOT NULL, OperationsType_ID INTEGER NOT NULL, Amount DOUBLE NOT NULL, EventDate DATE);"
	ddlBefore[6] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_Account FOREIGN KEY (Account_ID) REFERENCES Accounts(Account_ID);"
	ddlBefore[7] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_OperationsType FOREIGN KEY (OperationsType_ID) REFERENCES OperationsTypes(OperationsType_ID);"
	ddlBefore[8] = "ALTER TABLE Accounts ADD CONSTRAINT Accounts_UN UNIQUE KEY (Document_Number);"
	ddlBefore[9] = "INSERT INTO Accounts (Document_Number) VALUES ('123456');"
	ddlBefore[10] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (1, 'COMPRA A VISTA','DEBIT');"
	ddlBefore[11] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (2, 'COMPRA PARCELADA','DEBIT');"
	ddlBefore[12] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (3, 'SAQUE','DEBIT');"
	ddlBefore[13] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (4, 'PAGAMENTO','CREDIT');"

	var ddlAfter [3]string
	ddlAfter[0] = "Delete from Transactions;"
	ddlAfter[1] = "Delete from OperationsTypes;"
	ddlAfter[2] = "Delete from Accounts;"

	var database string
	var dbConnection string

	if len(os.Getenv("API_DB_DATABASE")) > 0 {
		database = os.Getenv("API_DB_DATABASE")
		dbConnection = os.Getenv("API_DB_CONNECTION")
	} else {
		database = dataBase
		dbConnection = connection
	}

	conn, err := sql.Open(database, dbConnection)
	if err != nil{
		panic(err)
	}

	return &DbTestConfig{
		Conn: conn,
		ddlBefore: ddlBefore,
		ddlAfter: ddlAfter,
	}
}

func (d DbTestConfig) Before(){
	for i := 0; i < len(d.ddlBefore); i++ {
		d.Conn.Exec(d.ddlBefore[i])
	}
}

func (d DbTestConfig) After() {
	for i := 0; i < len(d.ddlAfter); i++ {
		d.Conn.Exec(d.ddlAfter[i])
	}
	d.Conn.Close()
}
