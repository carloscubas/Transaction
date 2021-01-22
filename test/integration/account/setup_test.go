package account

import (
	"database/sql"
	"log"
	"transaction/internal/account"

	"github.com/gin-gonic/gin"
	_ "github.com/jmrobles/h2go"
	"go.uber.org/zap"
)

func main() {
	setupServer().Run()
}

// The engine with all endpoints is now extracted from the main function
func setupServer() *gin.Engine {

	conn := before()
	repository, _ := account.NewRepository(conn)

	handler := account.NewHandler(account.NewService(repository), zap.NewNop())

	router := gin.Default()

	api := router.Group("/v1")
	{
		api.POST("/transaction", handler.NewTransaction)
		api.POST("/accounts", handler.NewAccounts)
		api.GET("/accounts/:accountID", handler.GetAccounts)
	}

	return router
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
		_, err := conn.Exec(ddl[i])
		if err != nil {
			log.Fatalf("Can't exec ddl commands: %s", err)
			panic(err)
		}
	}
	return conn
}
