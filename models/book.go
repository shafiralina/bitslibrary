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

func (book *Book) Create() map[string]interface{} {

	if resp, ok := book.Validate(); !ok {
		return resp
	}

	GetDB().Create(book)
	resp := u.Message(true, "success")
	resp["book"] = book

	defer db.Close()
	return resp
}

func GetBook(id string) *Book {
	book := &Book{}
	err := GetDB().Where("id=?", id).First(book).Error
	if err != nil {
		return nil
	}

	defer db.Close()
	return book
}

func GetAllBook() []*Book {
	book := make([]*Book, 0)
	err := GetDB().Find(&book).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()
	return book
}

func (book *Book) Update(conn *gorm.DB, id string) map[string]interface{} {
	book1 := &Book{}
	conn.Debug().Where("id=?", id).First(&book1)
	book1 = book
	conn.Debug().Model(&book1).Updates(book)
	//t := time.Now()
	//GetDB().Table("books").Where("id=?", id).Update("updated_at", t)

	resp := u.Message(true, "success")
	resp["book"] = book1

	return resp
}

func Newest() []*Book {
	book := make([]*Book, 0)
	GetDB().Order("created_at desc").Limit(3).Find(&book)

	defer db.Close()
	return book
}

func Popular() []*Book {
	bookA := make([]*Book, 0)

	GetDB().Debug().Raw("SELECT a.* FROM books a ORDER BY coalesce((SELECT COUNT(book_id) FROM borrowds WHERE book_id = a.id GROUP BY book_id),0) deSC").Find(&bookA)
	//cari di borrowds dulu, trus baru di join sama books
	defer db.Close()
	return bookA
}
