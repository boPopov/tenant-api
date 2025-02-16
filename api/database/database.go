package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boPopov/tenant-api/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/**
 * Function ConnectDB.
 * Description: This function is used for establishing a connection with the database.
 */
func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	) //Setting up the connection string to the database.

	for {
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, _ := DB.DB()
			err = sqlDB.Ping() // Ensure database is actually reachable
			if err == nil {
				break
			}
		}

		log.Println("Database not ready. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	log.Println("Database connection established")

	//Executing AutoMigrations.
	if err := DB.AutoMigrate(&models.Tenant{}); err != nil {
		log.Fatal("Error was encounter while migrating the Tenant model. Error: ", err)
	}
}
