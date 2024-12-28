package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/tumenbayev/go-products-api/internal/controller/http/v1"
	"github.com/tumenbayev/go-products-api/internal/usecase/repository"
	"github.com/tumenbayev/go-products-api/internal/usecase"
	"github.com/tumenbayev/go-products-api/pkg/logger"
	"github.com/tumenbayev/go-products-api/pkg/httpserver"
	"github.com/tumenbayev/go-products-api/config"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	productRepo := repository.NewInMemoryProductRepository()
	productUseCase := usecase.NewProductUseCase(productRepo)

	handler := gin.New()
	v1.NewRouter(handler, l, productUseCase)
	httpServer := httpserver.New(handler, httpserver.WithPort(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
		case s := <-interrupt:
			l.Info("app - Run - signal: " + s.String())
		case err := <-httpServer.Notify():
			l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	if err := httpServer.Shutdown(); err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}