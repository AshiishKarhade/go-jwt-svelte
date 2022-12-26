package database

import (
	"github.com/AshiishKarhade/jwt-go-react/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "auth_user:Auth123@tcp(localhost:3306)/auth"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database")
	}
	DB = connection
	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic("Auto-migration failed")
	}
}
