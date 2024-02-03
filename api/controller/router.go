package controller

import (
	"student/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.POST("/student", CreateStudent)
	r.GET("/student", ReadAllStudents)
	r.GET("/student/:student_id", ReadStudentByID)
	r.PUT("/student/:student_id", UpdateStudent)
	r.DELETE("/student/:student_id", DeleteStudent)
	return r
}
