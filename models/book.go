package models

import (
	u "bitslibrary/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Book struct {
	Id        int64     `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Isbn      string    `json:"isbn"`
	Isbn13    string    `json:"isbn13"`
	Genre     string    `json:"genre"`
	Language  string    `json:"language"`
	DatePub   string    `json:"date_pub"`
	Pages     string    `json:"pages"`
	Sinopsis  string    `json:"sinopsis"`
	Price     float64   `json:"price"`
	Fineamt   float64   `json:"denda"`
}

func (Book) TableName() string {
	return "books"
}

func (book *Book) Validate() (map[string]interface{}, bool) {

	if book.Name == "" {
		return u.Message(false, "Book name must be written"), false
	}

	return u.Message(true, "success"), true
}

func (book *Book) Create(conn *gorm.DB) map[string]interface{} {

	if resp, ok := book.Validate(); !ok {
		return resp
	}

	conn.Create(book)
	resp := u.Message(true, "success")
	resp["book"] = book

	return resp
}

func GetBook(conn *gorm.DB, id string) *Book {
	book := &Book{}
	err := conn.Where("id=?", id).First(book).Error
	if err != nil {
		return nil
	}
	return book
}

func GetAllBook(conn *gorm.DB) []*Book {
	book := make([]*Book, 0)
	err := conn.Find(&book).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return book
}

func (book *Book) Update(conn *gorm.DB, id string) map[string]interface{} {
	book1 := &Book{}
	conn.Where("id=?", id).First(&book1)
	conn.Model(&book1).Updates(book)

	resp := u.Message(true, "success")
	resp["book"] = book1

	return resp
}

func Newest(conn *gorm.DB) []*Book {
	book := make([]*Book, 0)
	conn.Order("created_at desc").Limit(3).Find(&book)
	return book
}

func Popular(conn *gorm.DB) []*Book {
	bookA := make([]*Book, 0)
	conn.Debug().Raw("SELECT a.* FROM books a ORDER BY coalesce((SELECT COUNT(book_id) FROM borrowds WHERE book_id = a.id GROUP BY book_id),0) deSC").Find(&bookA)
	//cari di borrowds dulu, trus baru di join sama books
	return bookA
}
