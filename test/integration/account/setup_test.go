package account

import (
	"database/sql"
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
	router.GET("/v1/operationtypes", handler.GetOperationsTypes)

	return router
}