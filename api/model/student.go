package model

import (
	"net/http"
	"student/database"

	"github.com/gin-gonic/gin"
)

type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

func CreateStudent(c *gin.Context) {
	db := database.GetDB()
	var student Student
	db.Create(&student)
	c.JSON(200, student)
}

func ReadAllStudents(c *gin.Context) {
	db := database.GetDB()
	students := []Student{}
	db.Find(&students)
	c.JSON(200, students)
}

func ReadStudentByID(c *gin.Context) {
	db := database.GetDB()
	var student Student
	id := c.Params.ByName("student_id")
	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, student)
	}
}

func UpdateStudent(c *gin.Context) {
	db := database.GetDB()
	var student Student
	id := c.Params.ByName("student_id")

	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	var newStudent Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStudent.StudentID = student.StudentID

	if err := db.Model(&student).Where("student_id = ?", id).Save(&newStudent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, newStudent)
}

func DeleteStudent(c *gin.Context) {
	db := database.GetDB()
	var student Student
	id := c.Params.ByName("student_id")

	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Where("student_id = ?", id).Delete(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, "Deleted!")
}
