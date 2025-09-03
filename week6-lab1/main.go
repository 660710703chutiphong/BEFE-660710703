// main.go
package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"slices"
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
func getStudent(c *gin.Context) {
	id := c.Param("id")

	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}
func createStudent(c *gin.Context) {
	var newStudent Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newStudent.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	if newStudent.Year < 1 || newStudent.Year > 4{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Year must be between 1 and 4"})
		return
	}
	newStudent.ID = fmt.Sprintf("%d", len(students)+1)

	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)

}
func updateStudent(c *gin.Context) {
	id := c.Param("id")
	var updataStudent Student

	if err := c.ShouldBindJSON(&updataStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, student := range students {
		if student.ID == id {
			updataStudent.ID = id 
			students[i] = updataStudent
			c.JSON(http.StatusOK, updataStudent)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}
func deleteStudent(c *gin.Context) {
	id := c.Param("id")

	for i, student := range students{
		if student.ID == id{
			students = slices.Delete(students, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted Success"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Id not found"})
}
func main() {
	r := gin.Default()
	r.GET("/Heathly", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/students", getStudents)
		api.GET("/students/:id", getStudent)
		api.POST("/students", createStudent)
		api.PUT("/students/:id", updateStudent)
		api.DELETE("/students/:id", deleteStudent)
	}

	r.Run(":8080")
}
