package database

import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm" 
)

func Connect() (*gorm.DB, error) {
	dbConnectionString := "host=localhost user=lookana password=db123 dbname=lookana_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}