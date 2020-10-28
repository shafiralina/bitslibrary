package models

import (
	u "bitslibrary/utils"
	"time"
)

type Stock struct {
	Id        int64     `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	Qty       uint      `json:"quantity"`
	BookId    uint      `json:"book_id"`
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

func (stock *Stock) Create() map[string]interface{} {

	if resp, ok := stock.Validate(); !ok {
		return resp
	}

	GetDB().Create(stock)

	resp := u.Message(true, "success")
	resp["stock"] = stock

	defer db.Close()
	return resp
}
