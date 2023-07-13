package config

import (
	"github.com/mhdianrush/go-token-auth-jwt-mux/entities"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	logger := logrus.New()

	db, err := gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/go_token_auth_jwt_mux?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entities.User{})

	logger.Println("Database Connected")

	DB = db
}
