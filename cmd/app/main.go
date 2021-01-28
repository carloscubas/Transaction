package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
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

	defer db.Close()

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

	gracefullShutdown(router, sc.Server.Address)
}

func gracefullShutdown(router *gin.Engine, adress string) {
	server := &http.Server{
		Addr:    adress,
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}

func dataMigration(database string, connection string) {

	log.Info(fmt.Sprintf("Data Migration Start: %s://%s", database, connection))

	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("%s://%s", database, connection),
	)
	m.Steps(2)

	if err != nil {
		panic(err)
	}

	log.Info("Data Migration Finished")

}
