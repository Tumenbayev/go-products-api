package main

import (
	"log"

	"github.com/tumenbayev/go-products-api/config"
	"github.com/tumenbayev/go-products-api/internal/app"
	_ "github.com/tumenbayev/go-products-api/docs"
)

// @title           Products API
// @version         1.0
// @description     A Products API service
// @host           localhost:8080
// @BasePath       /v1
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}