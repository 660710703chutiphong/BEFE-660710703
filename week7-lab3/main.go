package main

import (
	"fmt"
	"os"
	"database/sql"
	// "github.com/lib/pq"
	"log"
)
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	} 
	return defaultValue
}
var db * sql.DB;

func initDB() {
	var err error

	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")
	conSt := fmt.Sprintf("Host=%s Name=%s User=%s Password=%s Port=%s, sslmode=disable", 
	host, name, user, password, port)
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
}