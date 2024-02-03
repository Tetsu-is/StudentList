package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDB := os.Getenv("MYSQL_DATABASE")
	dsn := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:password@tcp(mysql:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}
