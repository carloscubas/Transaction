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
func NewAPI(config Config, router *gin.Engine, repository Repository) error {

	notificationsHandler := NewHandler(NewService(repository), config.Logger)
	SetRoutes(notificationsHandler, config, router)
	return nil
}

func SetRoutes(handler *Handler, config Config, router *gin.Engine) {
	api := router.Group("/v1")
	{
		api.POST("/transaction", handler.NewTransaction)
		api.POST("/accounts", handler.NewAccounts)
		api.GET("/accounts/:accountID", handler.GetAccounts)
	}

	router.Run(config.Address)
}
