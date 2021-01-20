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

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	dataMigration(sc.Mysql.Connection)

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

func dataMigration(connection string){

	log.Info("Data Migration Start ")

	m, err := migrate.New(
		"file://db/migrations",
		"mysql://user:password@tcp(127.0.0.1:3306)/db")
	m.Steps(2)

	if err != nil {
		panic(err)
	}

	// 		"github://mattes:personal-access-token@mattes/migrate_test",

	log.Info("Data Migration Finished")
}