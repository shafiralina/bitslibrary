package models

import (
	u "bitslibrary/utils"
	"strconv"
)

type BorrowAPI struct {
	Borrow  Borrow    `json:"borrow"`
	Borrowd []Borrowd `json:"borrowd"`
}

func (borrowapi *BorrowAPI) Validate() (map[string]interface{}, bool) {
	//Get stock by book_id

	for _, x := range borrowapi.Borrowd {
		type ResultStock struct {
			Qty int
		}
		var resultstock ResultStock
		GetDB().Table("stocks").Select("qty").Where("book_id=?", x.BookId).Scan(&resultstock)

		if resultstock.Qty < x.Qty {
			return u.Message(false, "Books are out of stock!"), false
		}
	}
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

		x := Borrowd{BorrowId: borrow.ID, BookId: x.BookId, Qty: x.Qty, Subtotal: x.Subtotal, Price: x.Price}
		GetDB().Table("borrowds").Create(&x)
	}
	resp := u.Message(true, "success")
	resp["borrowapi"] = borrowapi
	return resp
}

//func Return(id string) map[string]interface{} {
//
//	GetDB().Table("borrows").Where("id=?", id).Update("status","F")
//	t := time.Now()
//	GetDB().Table("books").Where("id=?", id).Update("updated_at", t)
//
//	var result string
//	now := time.Now()
//	from,_ := time.Parse("2006-01-02", result)
//	days := now.Sub(from) / (24 * time.Hour)
//
//}
