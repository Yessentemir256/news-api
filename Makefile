# Название бинарного файла
BINARY_NAME=news-api

# Переменные окружения
ENV_FILE=.env
DB_CONTAINER_NAME=news-api-db

# Команды
build:
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) ./cmd/server

run: build
	@echo "Running the application..."
	@./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test ./...

lint:
	@echo "Running linters..."
	@golangci-lint run

docker-up:
	@echo "Starting Docker containers..."
	@docker-compose up -d

docker-down:
	@echo "Stopping Docker containers..."
	@docker-compose down

docker-logs:
	@echo "Viewing Docker logs..."
	@docker-compose logs -f

migrate-up:
	@echo "Applying database migrations..."
	@go run github.com/golang-migrate/migrate/v4/cmd/migrate -database "postgres://user:password@localhost:5432/news_db?sslmode=disable" -path ./migrations up

migrate-down:
	@echo "Reverting database migrations..."
	@go run github.com/golang-migrate/migrate/v4/cmd/migrate -database "postgres://user:password@localhost:5432/news_db?sslmode=disable" -path ./migrations down

.PHONY: build run clean test lint docker-up docker-down docker-logs migrate-up migrate-down
