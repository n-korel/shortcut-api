package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/n-korel/shortcut-api/internal/link"
	"github.com/n-korel/shortcut-api/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}, &user.User{})
}