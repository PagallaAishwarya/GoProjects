package models

import (
	"BookProject/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
// Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//Roles Struct
type Roles struct {
	gorm.Model
	RoleName string
}

//Accounts struct
type Accounts struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     Roles
}

//Books struct
type Books struct {
	gorm.Model
	user    Accounts
	Title   string
	Content string
}

func  GetAllBooks() []Books {
	var Books []Books
	db.Find(&Books)
	return Books
}

//CreateBook -creats a new book record
func (b *Books) CreateBook() *Books {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
// TableName Roles table name to be `user_accounts`
func (Roles) TableName() string {
	return "user_roles"
}

// TableName Accounts User table name to be `user_accounts`
func (Accounts) TableName() string {
	return "user_accounts"
}

// TableName Accounts User table name to be `user_accounts`
func (Books) TableName() string {
	return "books"
}
func init() {
	config.DBConnector()
	db = config.GetDB()
	db.AutoMigrate(&Roles{})
	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Books{})
}
