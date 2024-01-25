package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	StudentID uint   `json:"student_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

func main() {

	dsn := "tetsuro:11quin17@tcp(127.0.0.1:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Student{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	//create student
	r.POST("/student", func(c *gin.Context) {
		var student Student
		c.BindJSON(&student)
		db.Create(&student)
		c.JSON(200, student)
	})

	//read all students
	r.GET("/students", func(c *gin.Context) {
		var students []Student
		if err := db.Find(&students).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, students)
		}
	})

	//read student by student_id
	r.GET("/student/:student_id", func(c *gin.Context) {
		var student Student
		id := c.Params.ByName("student_id")
		if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, student)
		}
	})

	//update student by id
	r.PUT("/student/:student_id", func(c *gin.Context) {
		var student Student
		id := c.Params.ByName("student_id")
		if err := db.Where("student_id = ?", id).First(&student).Error; err != nil {
			c.AbortWithStatus(404)
		}
		c.BindJSON(&student)
		db.Save(&student)
		c.JSON(200, student)
	})

	//delete student by id

	r.Run()
}
