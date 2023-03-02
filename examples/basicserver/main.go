package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-kit/log"
	cfg "github.com/mailio/go-web3-kit/config"
	"github.com/mailio/go-web3-kit/examples/basicserver/api"
	"github.com/mailio/go-web3-kit/examples/basicserver/docs"
	w3srv "github.com/mailio/go-web3-kit/gingonic"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Conf global config
var Conf Config

// global Log
var Logger log.Logger

type Config struct {
	cfg.YamlConfig `yaml:",inline"`
}

func init() {
	w := log.NewSyncWriter(os.Stderr)
	Logger = log.NewLogfmtLogger(w)
	Logger.Log("question", "what is the meaning of life?", "answer", 42)
}

// @title Web3 Go Kit basic server
// @version 1.0
// @description This is a basic server example using go-web3-kit

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	var (
		configFile string
	)

	// configuration file optional path. Default:  current dir with  filename conf.yaml
	flag.StringVar(&configFile, "c", "conf.yaml", "Configuration file path.")
	flag.StringVar(&configFile, "config", "conf.yaml", "Configuration file path.")
	flag.Usage = usage
	flag.Parse()

	// var conf g.Config
	err := cfg.NewYamlConfig(configFile, &Conf)
	if err != nil {
		Logger.Log(err, "conf.yaml failed to load")
		panic("Failed to load conf.yaml")
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Go Web3 Kit API"
	docs.SwaggerInfo.Description = "This is a basic server example using go-web3-kit"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", Conf.Host, Conf.Port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{Conf.Scheme}
	ginSwagger.DefaultModelsExpandDepth(1)

	// server wait to shutdown monitoring channels
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	// init routing (for endpoints)
	router := w3srv.NewAPIRouter(&Conf.YamlConfig)

	root := router.Group("/api")
	{

		root.GET("/pong", api.PingPongAPI)
	}

	// start server
	srv := w3srv.Start(&Conf.YamlConfig, router)
	// wait for server shutdown
	go w3srv.Shutdown(srv, quit, done)

	Logger.Log("Server is ready to handle requests at", Conf.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		Logger.Log("Could not listen on %s: %v\n", Conf.Port, err)
	}

	<-done

}

// usage will print out the flag options for the server.
func usage() {
	usageStr := `Usage: operator [options]
	Server Options:
	-c, --config <file>              Configuration file path
`
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
