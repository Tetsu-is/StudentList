package model

import (
	"student/database"
)

type Student struct {
	StudentID int    `json:"student_id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

func Create(student *Student) (Student, error) {
	db := database.GetDB()
	if err := db.Create(&student).Error; err != nil {
		return *student, err
	}
	return *student, nil
}

func ReadAll() ([]Student, error) {
	db := database.GetDB()
	var students []Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func ReadByID(student_id string) (Student, error) {
	db := database.GetDB()
	var student Student
	if err := db.Where("student_id = ?", student_id).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (data *Student) Update() error {
	db := database.GetDB()
	if err := db.Where("student_id = ?", data.StudentID).Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (data *Student) Delete() error {
	db := database.GetDB()
	if err := db.Where("student_id = ?", data.StudentID).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
