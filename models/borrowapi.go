package models

import (
	u "bitslibrary/utils"
	"strconv"
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
		type ResultStock struct {
			Qty int
		}
		var resultstock ResultStock
		GetDB().Table("stocks").Select("qty").Where("book_id=?", x.BookId).Scan(&resultstock)
		//bisa pake find, gausa select
		if resultstock.Qty < x.Qty {
			return u.Message(false, "Books are out of stock!"), false
		}
	}
	defer db.Close()
	return u.Message(true, "success"), true
}

func (borrowapi *BorrowAPI) Create() map[string]interface{} {
	if resp, ok := borrowapi.Validate(); !ok {
		return resp
	}

	borrow := borrowapi.Borrow

	GetDB().Table("borrows").Create(&borrow)

	for _, x := range borrowapi.Borrowd {
		//pengurangan stock
		type ResultStock struct {
			Qty int
		}

		var resultstock ResultStock
		bookidstring := strconv.Itoa(x.BookId)
		GetDB().Table("stocks").Select("qty").Where("book_id=?", bookidstring).Scan(&resultstock)

		var stockcurrently = resultstock.Qty - x.Qty
		GetDB().Table("stocks").Where("book_id=?", x.BookId).Update("qty", stockcurrently)
		GetDB().Table("stocks").Where("book_id=?", x.BookId).Update("updated_at", time.Now())

		x := Borrowd{BorrowId: borrow.Id, BookId: x.BookId, Qty: x.Qty, Subtotal: x.Subtotal, Price: x.Price}
		GetDB().Table("borrowds").Create(&x)
	}
	resp := u.Message(true, "success")
	resp["borrowapi"] = borrowapi
	defer db.Close()
	return resp
}

func Return(id string) *ReturnResponse {
	response := &ReturnResponse{}

	GetDB().Table("borrows").Where("id=?", id).Update("status", "F")
	t := time.Now()
	GetDB().Table("books").Where("id=?", id).Update("updated_at", t)

	type Result struct {
		EndDate string
	}

	var result Result
	GetDB().Table("borrows").Select("end_date").Where("id=?", id).Scan(&result)

	now := time.Now()
	from, _ := time.Parse("2006-01-02", result.EndDate)
	days := now.Sub(from) / (24 * time.Hour)

	var denda int
	var total = 0

	borrowd := make([]*Borrowd, 0)
	GetDB().Table("borrowds").Where("borrow_id=?", id).Find(&borrowd)
	for _, x := range borrowd {
		type Result struct {
			Fineamt int
		}

		var result Result
		GetDB().Table("books").Select("fineamt").Where("id=?", x.BookId).Scan(&result)

		denda = int(days) * result.Fineamt * x.Qty
		total = total + denda

		//penambahan stock
		type ResultStock struct {
			Qty int
		}

		var resultstock ResultStock
		bookidstring := strconv.Itoa(x.BookId)
		GetDB().Table("stocks").Select("qty").Where("book_id=?", bookidstring).Scan(&resultstock)

		var stockcurrently = resultstock.Qty + x.Qty
		GetDB().Table("stocks").Where("book_id=?", x.BookId).Update("qty", stockcurrently)

		GetDB().Table("stocks").Where("book_id=?", x.BookId).Update("updated_at", time.Now())
	}

	response.LateDays = int(days)
	response.Fineamt = total

	defer db.Close()
	return response
}
