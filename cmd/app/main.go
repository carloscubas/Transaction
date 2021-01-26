package main

import (
	"database/sql"
	"flag"
	"fmt"
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
		Address:      sc.Server.Address,
		DbConnection: sc.Db.Connection,
	}

	db, err := sql.Open(sc.Db.Database, sc.Db.Connection)
	if err != nil {
		panic(err)
	}

	dataMigration(sc.Db.Database, sc.Db.Connection)

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

	repo, err := account.NewRepository(db)
	if err != nil {
		panic(err)
	}

	account.NewAPI(config, router, repo)

}

func dataMigration(database string, connection string) {

	log.Info(fmt.Sprintf("Data Migration Start: %s://%s", database, connection))

	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("%s://%s", database, connection),
	)
	if err != nil {
		panic(err)
	}

	err = m.Steps(2)
	if err != nil {
		panic(err)
	}


	log.Info("Data Migration Finished")

}
