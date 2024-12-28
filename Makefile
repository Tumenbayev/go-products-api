.PHONY: build run clean swagger

# Build the application
build:
	go build -o bin/app cmd/app/main.go

# Run the application
run:
	go run cmd/app/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Generate swagger documentation
swagger:
	swag init -g cmd/app/main.go -o docs

# Build and run in one command
all: clean build run

# Build docker image
docker-build:
	docker build -t go-products-api -f docker/Dockerfile .

# Run docker container
docker-run:
	docker run -p 8080:8080 -v $(PWD)/config:/app/config go-products-api

# Run with docker-compose
docker-compose-up:
	docker-compose up --build -d

# Stop docker-compose
docker-compose-down:
	docker-compose down
