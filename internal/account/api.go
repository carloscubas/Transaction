package account

import (

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger
	Adress string
}

// NewAPI create API handler
func NewAPI(config Config, router *gin.Engine) error {

	notificationsHandler := NewHandler(NewService(config), config.Logger)
	SetRoutes(notificationsHandler, config, router)
	return nil
}

func SetRoutes(handler *Handler, config Config, router *gin.Engine) {
	api := router.Group("/v1")
	{
		api.POST("/transaction", handler.NewTransaction)
	}

	router.Run(config.Adress)
}
