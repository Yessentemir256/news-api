# News API

News API is a simple RESTful service built using Go, Fiber, and Reform for managing news articles and their categories. It uses PostgreSQL as the database and supports Docker for easy deployment.

## Features

- **Edit News:** Update a news article and its categories using a POST request.
- **List News:** Retrieve a paginated list of news articles and their categories using a GET request.
- **Configuration:** Environment variables managed via Viper.
- **Logging:** Middleware-based logging with Fiber.
- **Connection Pooling:** Database connection pooling with `database/sql`.
- **Docker Support:** Easy deployment with Docker.

## Endpoints

### POST `/edit/:id`
Edit an existing news article.

#### Request Body:
```json
{
  "Id": 64,
  "Title": "Updated Title",
  "Content": "Updated Content",
  "Categories": [1, 2, 3]
}
```

- If a field is not provided, it will not be updated.

#### Response:
```json
{
  "success": true
}
```

### GET `/list`
Retrieve a paginated list of news articles.

#### Query Parameters:
- `page` (optional, default: 1): Page number.
- `limit` (optional, default: 10): Number of items per page.

#### Response:
```json
{
  "success": true,
  "news": [
    {
      "Id": 64,
      "Title": "Lorem ipsum",
      "Content": "Dolor sit amet",
      "Categories": [1, 2, 3]
    },
    {
      "Id": 1,
      "Title": "First News",
      "Content": "Sample content",
      "Categories": [1]
    }
  ]
}
```

## Setup

### Prerequisites

- Go 1.20+
- PostgreSQL 12+
- Docker & Docker Compose (optional for containerized setup)

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/Yessentemir256/news-api.git
   cd news-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure environment variables:
   Create a `.env` file with the following content:
   ```env
   DATABASE_DSN=postgres://user:password@localhost:5432/news_db?sslmode=disable
   SERVER_ADDRESS=:8080
   ```

4. Generate models using Reform:
   ```bash
   go generate ./...
   ```

5. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

### Docker Deployment

1. Build and run the Docker containers:
   ```bash
   docker-compose up --build
   ```

2. Access the application at `http://localhost:8080`.

### Database Migration

Use any migration tool, such as `golang-migrate`, to apply database migrations. Example:
```bash
migrate -path ./migrations -database "$DATABASE_DSN" up
```

## Project Structure

```
news-api/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── database/
│   └── init.go
├── handlers/
│   └── news.go
├── models/
│   └── models.go
├── migrations/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Tests

Write tests for handlers and database operations:
```bash
go test ./...
```

## License

This project is licensed under the MIT License.

