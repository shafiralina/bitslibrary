package models

import (
	u "bitslibrary/utils"
	"fmt"
	"time"
)

type Book struct {
	Id        int64     `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
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

	GetDB().Table("books").Create(book)

	resp := u.Message(true, "success")
	resp["book"] = book
	return resp
}

func GetBook(id string) *Book {
	book := &Book{}
	err := GetDB().Table("books").Where("id=?", id).First(book).Error
	if err != nil {
		return nil
	}
	return book
}

func GetAllBook() []*Book {
	book := make([]*Book, 0)
	err := GetDB().Table("books").Find(&book).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return book
}

func (book *Book) Update(id string) map[string]interface{} {

	GetDB().Table("books").Where("id=?", id).Update(book)
	t := time.Now()
	GetDB().Table("books").Where("id=?", id).Update("updated_at", t)

	resp := u.Message(true, "success")
	resp["book"] = book
	return resp
}

func Newest() []*Book {
	book := make([]*Book, 0)
	GetDB().Table("books").Order("created_at desc").Limit(3).Find(&book)
	return book
}

func Popular() []*Book {
	bookA := make([]*Book, 0)

	GetDB().Debug().Raw("SELECT a.* FROM books a ORDER BY coalesce((SELECT COUNT(book_id) FROM borrowds WHERE book_id = a.id GROUP BY book_id),0) deSC").Find(&bookA)

	return bookA
}
