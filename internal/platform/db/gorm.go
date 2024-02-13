package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/elvinlari/docker-golang/internal/task/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB, err error) {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
		os.Getenv("DB_TIMEZONE"))

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func RunMigration(db *gorm.DB) error {
    // Run migration for the first domain struct
    if err := db.AutoMigrate(&domain.Task{}); err != nil {
        return err
    }

    // Run migration for the second domain struct


    return nil
}
