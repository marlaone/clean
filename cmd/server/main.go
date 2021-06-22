package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"marla.one/clean/apps/ping"
	"marla.one/clean/pkg/apps"
	"marla.one/clean/pkg/config"
	"marla.one/clean/pkg/database"
	mrhttp "marla.one/clean/pkg/server/http"
)

func init() {
	if err := config.Load(); err != nil {
		panic(err)
	}
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db := &database.DatabaseAdapater{}
	if err := db.Open(); err != nil {
		// panic(err)
	}

	router := mrhttp.HttpRouter{Logger: logger}
	r := router.GetDefaultRouter()

	appsSystem := apps.NewAppsSystem(db, r)
	appsSystem.Register(ping.NewPingApp())

	logger.Info(fmt.Sprintf("Listening on :%s", viper.GetString("server.port")))
	http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("server.port")), r)
}
