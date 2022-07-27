package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	// getStudents "e-learning/api/student"
)

type student struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Birthdate string `json:"birthdate"`
}

var students = []student{
	{ID: "1", Name: "Blue Train", Surname: "Coltrane", Birthdate: "2022/01/01"},
	{ID: "2", Name: "Jeru", Surname: "Mulligan", Birthdate: "2022/01/01"},
	{ID: "3", Name: "Sarah Vaughan and Clifford Brown", Surname: "Vaughan", Birthdate: "2022/01/01"},
}

func getStudents(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)
}

func main() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	router := gin.Default()
	router.GET("/students", getStudents)

	router.Run("localhost:8080")
}
