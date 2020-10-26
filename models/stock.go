package models

import (
	u "bitslibrary/utils"
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Qty    uint `json:"quantity"`
	BookId uint `json:"book_id"`
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

	GetDB().Table("stocks").Create(stock)

	resp := u.Message(true, "success")
	resp["stock"] = stock
	return resp
}
