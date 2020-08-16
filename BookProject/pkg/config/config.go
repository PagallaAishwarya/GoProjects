package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	d * gorm.DB
)
// DBConnector : establishes database connection
func DBConnector() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=project0_user dbname=project0 password=pro_password sslmode=disable")
	// Migrate the schema
	if err == nil {
		fmt.Println("db connected")
	}
	d = db
}
func GetDB() *gorm.DB {
	return d
}
