package models

import (
	u "bitslibrary/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type Stock struct {
	Id        int64     `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	Qty       int       `json:"quantity"`
	BookId    int       `json:"book_id"`
}

func (Stock) TableName() string {
	return "stocks"
}

func (stock *Stock) Validate() (map[string]interface{}, bool) {

	if stock.BookId == 0 {
		return u.Message(false, "Book Id must be written"), false
	}
	if stock.Qty == 0 {
		return u.Message(false, "Quantity must be written"), false
	}

	return u.Message(true, "success"), true
}

func (stock *Stock) Create(conn *gorm.DB) map[string]interface{} {

	if resp, ok := stock.Validate(); !ok {
		return resp
	}

	conn.Create(stock)

	resp := u.Message(true, "success")
	resp["stock"] = stock

	return resp
}

func GetCurrentlyStock(bookid int) int {
	var resultstock Stock
	GetDB().Table("stocks").Select("qty").Where("book_id=?", bookid).Scan(&resultstock)
	return resultstock.Qty
}

func CalculationStock(bookid int, qty int) int {
	return GetCurrentlyStock(bookid) + qty
}
