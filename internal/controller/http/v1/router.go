package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tumenbayev/go-products-api/pkg/logger"
	"github.com/tumenbayev/go-products-api/internal/usecase"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine, logger logger.Interface, productUseCase usecase.ProductUseCase) {
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := handler.Group("/")

	NewProductHandler(router, logger, productUseCase)
}