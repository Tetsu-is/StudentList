package main

import (
	"student/controller"
	"student/database"
)

type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

func main() {

	database.Init()
	r := controller.GetRouter()
	r.Run()

	// dsn := "root:password@tcp(mysql:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Pong",
	// 	})
	// })

	// //create student ok
	// r.POST("/student", func(c *gin.Context) {
	// 	var student Student
	// 	c.BindJSON(&student)
	// 	db.Create(&student)
	// 	c.JSON(200, student)
	// })

	// //read all students ok
	// r.GET("/student", func(c *gin.Context) {
	// 	students := []Student{}
	// 	db.Find(&students)
	// 	c.JSON(200, students)
	// })

	// //read student by student_id ok
	// r.GET("/student/:student_id", func(c *gin.Context) {
	// 	var student Student
	// 	id := c.Params.ByName("student_id")
	// 	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
	// 		c.AbortWithStatus(404)
	// 	} else {
	// 		c.JSON(200, student)
	// 	}
	// })

	// //update student by id
	// r.PUT("/student/:student_id", func(c *gin.Context) {
	// 	var student Student
	// 	id := c.Params.ByName("student_id")

	// 	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	// 		return
	// 	}
	// 	var newStudent Student
	// 	if err := c.ShouldBindJSON(&newStudent); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	newStudent.StudentID = student.StudentID

	// 	if err := db.Model(&student).Where("student_id = ?", id).Save(&newStudent).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(200, newStudent)
	// })

	// //delete student by id
	// r.DELETE("/student/:student_id", func(c *gin.Context) {
	// 	var student Student
	// 	id := c.Params.ByName("student_id")

	// 	if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	// 		return
	// 	}

	// 	if err := db.Where("student_id = ?", id).Delete(&student).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(200, "Deleted!")
	// })

	// r.Run()
}
