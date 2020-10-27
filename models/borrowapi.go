package models

import (
	u "bitslibrary/utils"
	"fmt"
	"strconv"
)

type BorrowAPI struct {
	Borrow  Borrow    `json:"borrow"`
	Borrowd []Borrowd `json:"borrowd"`
}

//func (borrowapi *BorrowAPI) Validate() (map[string]interface{}, bool) {
//	//Get stock by book_id
//
//	for _, x := range borrowapi.Borrowd {
//		type ResultStock struct {
//			Stock string
//		}
//		var resultstock ResultStock
//		GetDB().Table("stocks").Select("qty").Where("book_id=?",x.BookId).Scan(&resultstock)
//		fmt.Println("stock=", resultstock.Stock)
//		fmt.Println("qty=", x.Qty)
//		fmt.Println("book_id=", x.BookId)
//
//		stokbukuint, _ := strconv.Atoi(resultstock.Stock)
//		if stokbukuint < x.Qty {
//			return u.Message(false, "Books are out of stock!"), false
//		}
//	}
//	return u.Message(true, "success"), true
//}

func (borrowapi *BorrowAPI) Create() map[string]interface{} {
	//if resp, ok := borrowapi.Validate(); !ok {
	//	return resp
	//}

	borrow := borrowapi.Borrow

	GetDB().Table("borrows").Create(&borrow)

	for _, x := range borrowapi.Borrowd {
		//pengurangan stock
		type ResultStock struct {
			Stock string
		}

		var resultstock ResultStock
		bookidstring := strconv.Itoa(x.BookId)
		GetDB().Table("stocks").Select("qty").Where("book_id=?", bookidstring).Scan(&resultstock)

		fmt.Println("INI STOCK =", resultstock.Stock)
		fmt.Println("ini id book =", bookidstring)

		stokbukuint, _ := strconv.Atoi(resultstock.Stock)
		var stockcurrently = stokbukuint - x.Qty
		GetDB().Table("stocks").Where("book_id=?", x.BookId).Update("qty", stockcurrently)

		x := Borrowd{BorrowId: borrow.ID, BookId: x.BookId, Qty: x.Qty, Subtotal: x.Subtotal, Price: x.Price}
		GetDB().Table("borrowds").Create(&x)
	}
	resp := u.Message(true, "success")
	resp["borrowapi"] = borrowapi
	return resp
}
