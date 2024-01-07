package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password= dbname=schalter port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        return nil, err
    }

    return db, nil
}