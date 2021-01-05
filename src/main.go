package main

import (
	"errors"
	"flag"

	"github.com/NODO-UH/master-cut/src/api"
	"github.com/NODO-UH/master-cut/src/conf"
	"github.com/gin-gonic/gin"
)

var confPath *string

func init() {
	confPath = flag.String("conf", "config.json", "path to configuration file")
}

func main() {
	flag.Parse()

	// Setup configuration
	if confPath == nil {
		panic(errors.New("config file unknown"))
	}
	if err := conf.SetupConfiguration(*confPath); err != nil {
		panic(err)
	}

	// Start Gin server REST API
	server := gin.New()
	server.Use(gin.Logger())

	server.POST("/cut", api.Cut)
	server.POST("/uncut", api.Uncut)

	// Run in port saved in PORT environment variable
	if err := server.Run(); err != nil {
		panic(err)
	}
}
