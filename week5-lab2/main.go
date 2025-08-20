package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUsers(c *gin.Context) {
	User := []User{{ID: "3", Name: "Chutiphong"}}
	c.JSON(200, User)
}
func main() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.Run(":8080")
}
