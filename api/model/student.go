package model

import (
	"student/database"
)

type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

func Create(student *Student) {
	db := database.GetDB()
	db.Create(&student)
}

func ReadAll() []Student {
	//(students []Student)でも良い？？
	db := database.GetDB()
	var students []Student
	db.Find(&students)
	return students
}

func ReadByID(student_id string) Student {
	db := database.GetDB()
	var student Student
	db.Where("student_id = ?", student_id).First(&student)
	return student
}

func (data *Student) Update() {
	db := database.GetDB()
	db.Where("student_id = ?", data.StudentID).Save(&data)
}

func (data *Student) Delete() {
	db := database.GetDB()
	db.Where("student_id = ?", data.StudentID).Delete(&data)
}
