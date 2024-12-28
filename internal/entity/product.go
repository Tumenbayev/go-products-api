package entity

// @Description Product represents a product in the system
type Product struct {
	// @Description Unique identifier for the product
	SKU     string  `json:"sku"`
	// @Description Name of the product
	Name    string  `json:"name"`
	// @Description Category of the product
	Category string `json:"category"`
	// @Description Price information of the product
	Price   Price   `json:"price"`
}

// @Description Price contains pricing information for a product
type Price struct {
	// @Description Original price in cents/pennies
	Original           int     `json:"original"`
	// @Description Final price after applying discounts in cents/pennies
	Final              int     `json:"final"`
	// @Description Discount percentage if applicable
	DiscountPercentage *string `json:"discount_percentage"`
	// @Description Currency code (e.g., EUR)
	Currency           string  `json:"currency"`
}