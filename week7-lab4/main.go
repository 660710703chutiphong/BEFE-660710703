package main

import (
	"fmt"
	"os"
	"database/sql"
	// "github.com/lib/pq"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	} 
	return defaultValue
}
var db *sql.DB

func initDB() {
	var err error

	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")
	conSt := fmt.Sprintf(
    "host=%s dbname=%s user=%s password=%s port=%s sslmode=disable",
    host, name, user, password, port,
)

	fmt.Println(conSt)
	//
	db, err := sql.Open("postgres", conSt)

	if err != nil {
		log.Fatal("Failed to open a Database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect Database")
	}
	log.Println("Sucss Connected to Database")
}
func main() {
	initDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context){
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealthy","error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message" : "healthy"})
	})

	// api := r.Group("/api/v1")
	// {
	// 	api.GET("/books", getBooks)
	// 	api.GET("/books/:id", getBook)
	// 	api.POST("/books", createBook)
	// 	api.PUT("/books/:id", updateBook)
	// 	api.DELETE("/books/:id", deleteBook)
	// }

	r.Run(":8080")
}