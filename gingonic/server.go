package gingonic

// Copyright 2021 Mailio All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mailio/go-web3/kit/config"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// NewAPIRouter initializes all public/secure api routes including short api description for swagger documentation
func NewAPIRouter(conf *config.YamlConfig) *gin.Engine {
	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Public API Definitions
	public := router.Group("/")

	// Optionally include swagger (from conf.yaml)
	if conf.Swagger {
		public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}

// Start - start the gin server
func Start(conf *config.YamlConfig, router *gin.Engine) *http.Server {
	//  starting server (Default With the Logger and Recovery middleware already attached)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: router,
	}
	return srv
}

// Shutdown gracefully shutdown of the server
func Shutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	fmt.Printf("Server is shutting down...")
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
