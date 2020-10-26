package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	//jadi kalo pertama run, dia create users
	//tapi pas panggil api, butuhnya user
	//tadi misal yg udh ke create itu tak rename jadi user
	//trus ku panggil api, dia baru bisa. coding bikin tablemana

	//gak tau, pokoknya base.go ini buat connect database
	//modelnya ya di user.go
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.AutoMigrate(&User{}, &Book{}, &Stock{}, &Borrow{}, &Borrowd{})
}

func GetDB() *gorm.DB {
	return db
}
