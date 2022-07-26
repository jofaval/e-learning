package student

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name      string
	Surname   string
	Birthdate string
}

var students = []Student{
	{Name: "Blue Train", Surname: "Coltrane", Birthdate: "2022/01/01"},
	{Name: "Jeru", Surname: "Mulligan", Birthdate: "2022/01/01"},
	{Name: "Sarah Vaughan and Clifford Brown", Surname: "Vaughan", Birthdate: "2022/01/01"},
}

func getStudents(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)
}
