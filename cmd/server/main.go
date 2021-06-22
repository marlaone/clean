package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"marla.one/clean/apps/ping"
	"marla.one/clean/pkg/apps"
	"marla.one/clean/pkg/config"
	"marla.one/clean/pkg/database"
	mrhttp "marla.one/clean/pkg/server/http"
)

var configName string

func init() {

}

func ServerCmd() {
	if err := config.Load(configName); err != nil {
		panic(err)
	}
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

func main() {

	var cmdServer = &cobra.Command{
		Use:   "start",
		Short: "http server",
		Long:  `start http server`,
		Run: func(cmd *cobra.Command, args []string) {
			ServerCmd()
		},
	}

	cmdServer.Flags().StringVarP(&configName, "config", "c", "dev.config", "The server config to load.")

	var serverCmd = &cobra.Command{Use: "server"}
	serverCmd.AddCommand(cmdServer)
	serverCmd.Execute()

}
