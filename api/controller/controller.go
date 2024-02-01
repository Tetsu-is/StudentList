package controller

import (
	"net/http"
	"student/model"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student model.Student
	c.BindJSON(&student)
	model.Create(&student)
	c.JSON(http.StatusOK, student)
}

func ReadAllStudents(c *gin.Context) {
	students := model.ReadAll()
	c.JSON(http.StatusOK, students)
}

func ReadStudentByID(c *gin.Context) {
	id := c.Params.ByName("student_id")
	student := model.ReadByID(id)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("student_id")
	data := model.ReadByID(id)

	var new model.Student
	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Name, data.Age = new.Name, new.Age

	data.Update()

	c.JSON(http.StatusOK, new)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("student_id")
	data := model.ReadByID(id)
	data.Delete()
	c.JSON(200, data)
}
