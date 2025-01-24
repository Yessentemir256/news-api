package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"      // Импортируем для работы с PostgreSQL
	"github.com/news-api/sqlc" // Импортируй свой пакет db, где сгенерированы методы sqlc
)

func main() {
	// Настройки подключения
	databaseURL := "postgres://username:password@localhost:5432/dbname"

	// Создание подключения к базе данных
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Создание экземпляра Queries
	queries := db.New(conn)

	// Пример использования сгенерированного кода

	// Создание новости
	err = createNews(queries)
	if err != nil {
		log.Fatalf("Error creating news: %v\n", err)
	}

	// Получение новости по ID
	news, err := getNewsById(queries, 1)
	if err != nil {
		log.Fatalf("Error getting news: %v\n", err)
	}
	fmt.Printf("News: %+v\n", news)

	// Добавление категории к новости
	err = addNewsCategory(queries, 1, 2)
	if err != nil {
		log.Fatalf("Error adding category: %v\n", err)
	}

	// Получение категорий по ID новости
	categories, err := getCategoriesByNewsId(queries, 1)
	if err != nil {
		log.Fatalf("Error getting categories: %v\n", err)
	}
	fmt.Printf("Categories: %+v\n", categories)
}

// Функция для создания новости
func createNews(queries *db.Queries) error {
	params := db.CreateNewsParams{
		Title:   "New Title",
		Content: "Content of the new news article",
	}
	return queries.CreateNews(context.Background(), params)
}

// Функция для получения новости по ID
func getNewsById(queries *db.Queries, id int32) (db.News, error) {
	return queries.GetNewsById(context.Background(), id)
}

// Функция для добавления категории к новости
func addNewsCategory(queries *db.Queries, newsId, categoryId int32) error {
	params := db.AddNewsCategoryParams{
		Newsid:     newsId,
		Categoryid: categoryId,
	}
	return queries.AddNewsCategory(context.Background(), params)
}

// Функция для получения категорий по ID новости
func getCategoriesByNewsId(queries *db.Queries, newsId int32) ([]int32, error) {
	return queries.GetCategoriesByNewsId(context.Background(), newsId)
}
