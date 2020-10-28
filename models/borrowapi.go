package models

import (
	u "bitslibrary/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type BorrowAPI struct {
	Borrow  Borrow    `json:"borrow"`
	Borrowd []Borrowd `json:"borrowd"`
}

type ReturnResponse struct {
	Fineamt  int `json:"fineamt"`
	LateDays int `json:"late_days"`
}

func (borrowapi *BorrowAPI) Validate() (map[string]interface{}, bool) {
	//Get stock by book_id

	for _, x := range borrowapi.Borrowd {
		stock := GetCurrentlyStock(x.BookId)
		if stock < x.Qty {
			return u.Message(false, "Books are out of stock!"), false
		}
	}
	return u.Message(true, "success"), true
}

func (borrowapi *BorrowAPI) Create(conn *gorm.DB) map[string]interface{} {
	if resp, ok := borrowapi.Validate(); !ok {
		return resp
	}

	borrow := borrowapi.Borrow

	conn.Table("borrows").Create(&borrow)

	for _, x := range borrowapi.Borrowd {
		//pengurangan stock
		//stock := GetCurrentlyStock(x.BookId)
		//var stockcurrently = stock - x.Qty

		var stockcurrently = CalculationStock(x.BookId, -x.Qty)
		stock1 := &Stock{}
		conn.Table("stocks").Where("book_id=?", x.BookId).First(&stock1)
		conn.Model(&stock1).Update("qty", stockcurrently)

		//create borrow
		x := Borrowd{BorrowId: borrow.Id, BookId: x.BookId, Qty: x.Qty, Subtotal: x.Subtotal, Price: x.Price}
		conn.Table("borrowds").Create(&x)
	}
	resp := u.Message(true, "success")
	resp["borrowapi"] = borrowapi

	return resp
}

func Return(conn *gorm.DB, id string) *ReturnResponse {
	response := &ReturnResponse{}

	borrow1 := &Borrow{}
	conn.Table("borrows").Where("id=?", id).First(&borrow1)
	conn.Model(&borrow1).Update("status", "F")

	var result Borrow
	conn.Table("borrows").Select("end_date").Where("id=?", id).Scan(&result)

	now := time.Now()
	from, _ := time.Parse("2006-01-02", result.EndDate)
	days := now.Sub(from) / (24 * time.Hour)

	var denda int
	var total = 0

	borrowd := make([]*Borrowd, 0)
	conn.Table("borrowds").Where("borrow_id=?", id).Find(&borrowd)
	for _, x := range borrowd {
		var result Book
		conn.Table("books").Select("fineamt").Where("id=?", x.BookId).Scan(&result)

		denda = int(days) * int(result.Fineamt) * x.Qty
		total = total + denda

		//penambahan stock
		//stock := CalculateStock(x.BookId)
		//var stockcurrently = stock + x.Qty

		var stockcurrently = CalculationStock(x.BookId, x.Qty)
		stock1 := &Stock{}
		conn.Table("stocks").Where("book_id=?", x.BookId).First(&stock1)
		conn.Model(&stock1).Update("qty", stockcurrently)
	}

	response.LateDays = int(days)
	response.Fineamt = total

	return response
}
