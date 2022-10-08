package database

import (
	"assignment2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "hacktiv8go"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo))
	if err != nil {
		fmt.Println("Error open connection to postgre", err)
		return
	}

	fmt.Println("Success open connection to postgre")

	err = db.Debug().AutoMigrate(models.Order{}, models.Item{})
	if err != nil {
		fmt.Println("Error when migrate DB postgre", err)
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
