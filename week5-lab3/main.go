// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student struct
type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var students = []Student{
	{ID: "1", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.50},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 3.75},
}

// GET /api/v1/students
func getStudents(c *gin.Context) {
	// ตรวจสอบว่ามี query parameter "year" หรือไม่
	yearQuery := c.Query("year")

	if yearQuery != "" {
		// Filter students by year
		var filtered []Student
		for _, student := range students {
			if fmt.Sprint(student.Year) == yearQuery {
				filtered = append(filtered, student)
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}

	// Return all students
	c.JSON(http.StatusOK, students)
}
func main() {
	r := gin.Default()
	r.GET("/Heathly", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/students", getStudents)
	}

	r.Run(":8080")
}
