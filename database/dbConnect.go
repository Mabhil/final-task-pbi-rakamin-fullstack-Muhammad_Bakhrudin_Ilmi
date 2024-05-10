package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	var loginDB = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", DB_HOST, DB_PORT, DB_USERNAME, DB_NAME, DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(loginDB), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	for _, m := range MigrationDB() {
		err := db.AutoMigrate(m.Migration)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Database migrated!")

	fmt.Println("Database connected!")
	DB = db
}
