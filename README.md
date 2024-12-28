# go-products-api

## Overview

This project is a REST API that applies discounts to products and supports filtering. It is built with Go, because as I heard during the first call, you're using Go in your company.

## Features

- **Discount Application**: 
  - 30% off for products in the boots category.
  - 15% off for the product with SKU 000003.
  - The largest discount is applied when multiple discounts are applicable.
- **Product Filtering**:
  - Filter by category.
  - Filter by price less than a specified amount (optional).
- **API Endpoints**:
  - `GET /products`: Returns a list of products with discounts applied, filtered by category and/or price.

## Build and Run

### Prerequisites

- Docker
- Docker Compose

### Build and Run in One Command

```bash
make all
```

### Run with Docker Compose

```bash
make docker-compose-up
```

### Stop Docker Compose

```bash
make docker-compose-down
```

## Testing

Tests are included and can be run with a single command. Ensure you have Go installed.

```bash
go test ./...
```
or it can be run directly in docker container


## Architecture and Design Decisions

- **Language and Framework**: The API is built using Go for its performance and concurrency capabilities, which are essential for handling large datasets.
- **Data Storage**: Products are stored in-memory for simplicity and performance. This can be easily adapted to use a database if needed.
- **Docker**: The application and its dependencies are containerized using Docker, ensuring consistent environments across different machines.
- **Scalability**: The API is designed by following clean code principles and following Domain Driven Design, which makes it easy to scale and maintain.

## API Usage

### `GET /products`

- **Query Parameters**:
  - `category`: Filter by product category.
  - `priceLessThan`: Filter by price less than or equal to the specified value (pre-discount).

- **Response**: Returns a list of up to 5 products with discounts applied.

#### Example Request

```http
GET /products?category=boots&priceLessThan=80000
```

#### Example Response

```json
[
  {
    "sku": "000001",
    "name": "BV Lean leather ankle boots",
    "category": "boots",
    "price": {
      "original": 89000,
      "final": 62300,
      "discount_percentage": "30%",
      "currency": "EUR"
    }
  }
]
```

I didn't do overengineering, but I did follow clean code principles and domain driven design.
In future we could add more security on API level, caching, authentication, etc.

 **Performance Considerations**: 

Current implementation is quite simple, since we're not using real database and dataset is small, but it is easy to scale and maintain, we can switch from in-memory storage to database and add caching layer, or we could use elasticsearch to improve search performance.

**Future improvements:**

- As the product list grows, we may consider implementing pagination to limit the number of products returned in a single request. This will help manage response sizes and improve performance.
- We could add caching layer to improve performance of frequently accessed data.
- We could use elasticsearch to improve search performance.

