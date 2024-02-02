package controller

import (
	"net/http"
	"student/model"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student model.Student
	c.BindJSON(&student)
	if student.Name == "" || student.Age == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please fill in the name and age"})
		return
	}

	new, err := model.Create(&student)
	if err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, new)
}

func ReadAllStudents(c *gin.Context) {
	students, err := model.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func ReadStudentByID(c *gin.Context) {
	id := c.Params.ByName("student_id")
	student, err := model.ReadByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("student_id")
	data, err := model.ReadByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var new model.Student
	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new.StudentID = data.StudentID
	data.Name, data.Age = new.Name, new.Age

	if err := data.Update(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, new)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("student_id")
	data, err := model.ReadByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}
