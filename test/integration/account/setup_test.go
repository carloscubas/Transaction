package account

import (
	"database/sql"
	"log"
	"os"
	"transaction/internal/account"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmrobles/h2go"
	"go.uber.org/zap"
)


// The engine with all endpoints is now extracted from the main function
func setupServer(db *sql.DB) *gin.Engine {

	repository, _ := account.NewRepository(db)

	handler := account.NewHandler(account.NewService(repository), zap.NewNop())

	router := gin.Default()

	router.POST("/v1/transaction", handler.NewTransaction)
	router.POST("/v1/accounts", handler.NewAccounts)
	router.GET("/v1/accounts/:accountID", handler.GetAccounts)

	return router
}

func before() *sql.DB{

	var ddl [14]string

	ddl[0] = "DROP TABLE IF EXISTS Transactions;"
	ddl[1] = "DROP TABLE IF EXISTS OperationsTypes;"
	ddl[2] = "DROP TABLE IF EXISTS Accounts;"
	ddl[3] = "CREATE TABLE Accounts (Account_ID INTEGER auto_increment primary key, Document_Number varchar(100) NOT NULL);"
	ddl[4] = "CREATE TABLE OperationsTypes ( OperationsType_ID INTEGER auto_increment primary key, Description varchar(200) NOT NULL, OperationsType varchar(200) NOT NULL);"
	ddl[5] = "CREATE TABLE Transactions (Transaction_ID INTEGER auto_increment primary key, Account_ID INTEGER NOT NULL, OperationsType_ID INTEGER NOT NULL, Amount DOUBLE NOT NULL, EventDate DATE);"
	ddl[6] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_Account FOREIGN KEY (Account_ID) REFERENCES Accounts(Account_ID);"
	ddl[7] = "ALTER TABLE Transactions ADD CONSTRAINT Transactions_FK_OperationsType FOREIGN KEY (OperationsType_ID) REFERENCES OperationsTypes(OperationsType_ID);"
	ddl[8] = "ALTER TABLE Accounts ADD CONSTRAINT Accounts_UN UNIQUE KEY (Document_Number);"
	ddl[9] = "INSERT INTO Accounts (Document_Number) VALUES ('123456');"
	ddl[10] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (1, 'COMPRA A VISTA','DEBIT');"
	ddl[11] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (2, 'COMPRA PARCELADA','DEBIT');"
	ddl[12] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (3, 'SAQUE','DEBIT');"
	ddl[13] = "INSERT INTO OperationsTypes (OperationsType_ID, Description,OperationsType) VALUES (4, 'PAGAMENTO','CREDIT');"

	var database string
	var dbConnection string

	if len(os.Getenv("API_DB_DATABASE")) > 0 {
		database = os.Getenv("API_DB_DATABASE")
		dbConnection = os.Getenv("API_DB_CONNECTION")
	} else {
		//database = "h2"
		//dbConnection = "h2://sa@localhost/test?mem=true&logging=info"
		database = "mysql"
		dbConnection = "root:@tcp(127.0.0.1:3306)/dbtest"
	}

	conn, err := sql.Open(database, dbConnection)

	if err != nil {
		log.Fatalf("Can't connet to H2 Database: %s", err)
	}

	tx, _ := conn.Begin()
	for i := 0; i < len(ddl); i++ {
		log.Printf("Execute %s", ddl[i])
		response, err := tx.Exec(ddl[i])
		log.Printf("Response %v", response)
		if err != nil {
			log.Fatalf("Can't exec ddl commands: %s", err)
		}
	}
	tx.Commit()
	return conn
}

func after(conn *sql.DB){
	conn.Close()
}