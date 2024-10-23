package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Init() {
	var err error
	dsn := "postgres://postgres:root@localhost:5432/Users?sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully.")

	Db.AutoMigrate(Users{})

}

type Users struct {
	gorm.Model
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int16  `json:"age"`
}

func (Users) TableName() string {
	return "userinfo"
}
