package account

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Config struct {
	Logger       *zap.Logger
	Adress       string
	DbConnection string
}

// NewAPI create API handler
func NewAPI(config Config, router *gin.Engine) error {

	repo, err := NewMysqlRepository(config)
	if err != nil {
		return err
	}

	notificationsHandler := NewHandler(NewService(config, repo), config.Logger)
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
