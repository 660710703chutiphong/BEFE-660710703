package main

import (
	"fmt"
	"os"
)
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	} 
	return defaultValue
}
func main() {
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")
	
	conSt := fmt.Sprintf("Host=%s Name=%s User=%s Password=%s Port=%s", host, name, user, password, port)
	fmt.Println(conSt)
}