package config 

import (
	"fmt" 

	"os" 

	"go-gin-postgres-crud/models" 

	"github.com/joho/godotenv" 

	"gorm.io/driver/postgres" 

	"gorm.io/gorm" 
)

var DB *gorm.DB

func ConnectDatabase() {

	godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection failed")
	}

	DB = database

	DB.AutoMigrate(&models.Product{})
}