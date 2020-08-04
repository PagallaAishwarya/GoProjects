package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB
// DBConnector : establishes database connection
func DBConnector() {
	DB, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=project0_user dbname=project0 password=pro_password sslmode=disable")
	defer DB.Close()
	// Migrate the schema
	if err == nil {
		fmt.Println("db connected")
	}

	DB.AutoMigrate(&Roles{})
	DB.AutoMigrate(&Accounts{})
	DB.AutoMigrate(&Books{})
}
