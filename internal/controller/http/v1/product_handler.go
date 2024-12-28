package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tumenbayev/go-products-api/internal/entity"
	"github.com/tumenbayev/go-products-api/internal/usecase"
	"github.com/tumenbayev/go-products-api/pkg/logger"
)

type getProductsResponse struct {
	Products []entity.Product `json:"products"`
}

type response struct {
	Error string `json:"error"`
}

type ProductHandler struct {
	useCase usecase.ProductUseCase
	logger  logger.Interface
}

// NewProductHandler initializes the product handler and registers routes.
func NewProductHandler(router *gin.RouterGroup, logger logger.Interface, useCase usecase.ProductUseCase) {
	handler := &ProductHandler{
		useCase: useCase,
		logger:  logger,
	}
	router.GET("/products", handler.GetProducts)
}

// @Summary     Get Products
// @Description Retrieve a list of products with optional filters and applied discounts
// @ID          get-products
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       category query string false "Category to filter by (e.g., boots, sandals, sneakers)"
// @Param       priceLessThan query integer false "Maximum price in cents/pennies (e.g., 89000 for 890.00€)"
// @Success     200 {object} getProductsResponse "Successfully retrieved products"
// @Failure     400 {object} response "Invalid request parameters"
// @Failure     500 {object} response "Internal server error"
// @Router      /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	filter, err := parseProductFilter(c)
	if err != nil {
		h.logger.Error("GetProducts: invalid filter parameters", err)
		c.JSON(http.StatusBadRequest, response{Error: err.Error()})
		return
	}

	products, err := h.useCase.GetProducts(filter)
	if err != nil {
		h.logger.Error("GetProducts: failed to retrieve products", err)
		c.JSON(http.StatusInternalServerError, response{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, getProductsResponse{Products: products})
}

func parseProductFilter(c *gin.Context) (usecase.ProductFilter, error) {
	var filter usecase.ProductFilter

	if category := c.Query("category"); category != "" {
		filter.Category = category
	}

	if priceStr := c.Query("priceLessThan"); priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			return filter, fmt.Errorf("invalid priceLessThan parameter")
		}
		filter.PriceLessThan = price
	}

	return filter, nil
}