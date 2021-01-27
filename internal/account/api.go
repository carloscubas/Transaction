package account

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Config struct {
	Logger       *zap.Logger
	Address      string
	DbConnection string
}

// NewAPI create API handler
func NewAPI(config Config, router *gin.Engine, repository Repository, log chan string) error {

	notificationsHandler := NewHandler(NewService(repository), config.Logger, log)
	SetRoutes(notificationsHandler, config, router)
	return nil
}

// SetRoutes is used to aggregate all user endpoints.
func SetRoutes(handler *Handler, config Config, router *gin.Engine) {
	api := router.Group("/v1")
	{
		api.POST("/transaction", handler.NewTransaction)
		api.POST("/accounts", handler.NewAccounts)
		api.GET("/accounts/:accountID", handler.GetAccounts)
		api.GET("/operationtypes", handler.GetOperationsTypes)
	}

	router.Run(config.Address)
}
