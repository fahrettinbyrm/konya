package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
