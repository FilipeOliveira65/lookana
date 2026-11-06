package database 

import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm" 

	"lookana/api/internal/models"
)

func CreateNewService() (db gorm.DB, service models.ServiceData) {
}
