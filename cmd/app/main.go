package main

import (
	"flag"
	"net/http"
	"time"
	"transaction/internal/account"
	"transaction/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/util"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {

	var configFile string

	flag.StringVar(&configFile, "c", "configs/dev.yaml", "config file path")
	flag.Parse()

	sc, err := config.LoadServiceConfig(configFile)
	if err != nil {
		log.Fatalf("main: could not load service configuration [%v]", err)
	}

	config := account.Config{
		Logger:       zap.NewNop(),
		Adress:       sc.Server.Address,
		DbConnection: sc.Mysql.Connection,
	}

	gin.SetMode(sc.Server.Mode)
	router := gin.New()
	router.Use(JSONLogMiddleware())
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	account.NewAPI(config, router)

}

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := util.GetDurationInMillseconds(start)

		entry := log.WithFields(log.Fields{
			"client_ip":  util.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
