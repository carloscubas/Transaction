package main

import (
	"flag"
	"net/http"
	"transaction/internal/account"
	"transaction/internal/config"

	httpLogger "github.com/elafarge/gin-http-logger"
	"github.com/gin-gonic/gin"
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


	httpLoggerConf := httpLogger.AccessLoggerConfig{
		LogrusLogger:   log.StandardLogger(),
		BodyLogPolicy:  httpLogger.LogBodiesOnErrors,
		MaxBodyLogSize: 100,
		DropSize:       5,
		RetryInterval:  5,
	}

	router.Use(httpLogger.New(httpLoggerConf))

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	account.NewAPI(config, router)

}

/*
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

 */
