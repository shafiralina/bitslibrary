package models

import (
	u "bitslibrary/utils"
	"github.com/jinzhu/gorm"
)

type Borrowd struct {
	gorm.Model
	BorrowId uint    `json:"borrow_id"`
	BookId   uint    `json:"book_id"`
	Qty      uint    `json:"quantity"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}

func (borrowd *Borrowd) Validate() (map[string]interface{}, bool) {
	//Get stock by book_id
	type ResultStock struct {
		Stock uint
	}
	var resultstock ResultStock
	GetDB().Raw("SELECT qty FROM stocks WHERE book_id = ?", borrowd.BookId).Scan(&resultstock)
	if resultstock.Stock < borrowd.Qty {
		return u.Message(false, "Books are out of stock!"), false
	}
	if borrowd.BorrowId == 0 {
		return u.Message(false, "BorrowId must be written"), false
	}
	if borrowd.BookId == 0 {
		return u.Message(false, "BookId must be written"), false
	}

	return u.Message(true, "success"), true
}

func (borrowd *Borrowd) Create() map[string]interface{} {
	if resp, ok := borrowd.Validate(); !ok {
		return resp
	}

	//Get price by book_id
	type ResultPrice struct {
		Price float64
	}
	var resultprice ResultPrice
	GetDB().Raw("SELECT price FROM books WHERE id = ?", borrowd.Price).Scan(&resultprice)

	//Subtotal
	var subtotal = borrowd.Price * float64(borrowd.Qty)

	//Create borrowd
	data := Borrowd{BorrowId: borrowd.BorrowId, BookId: borrowd.BookId, Qty: borrowd.Qty, Price: resultprice.Price, Subtotal: subtotal}
	GetDB().Table("borrowds").Create(&data)

	//Update stock
	type ResultStock struct {
		Stock uint
	}
	var resultstock ResultStock
	GetDB().Raw("SELECT qty FROM stocks WHERE book_id = ?", borrowd.BookId).Scan(&resultstock)

	var stockcurrently = resultstock.Stock - borrowd.Qty
	GetDB().Table("stocks").Where("book_id=?", borrowd.BookId).Update("stock", stockcurrently)

	resp := u.Message(true, "success")
	resp["borrowd"] = borrowd
	return resp
}
