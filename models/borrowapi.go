package models

import (
	u "bitslibrary/utils"
	"time"
)

type BorrowAPI struct {
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	UserId      uint    `json:"user_id"`
	Status      string  `json:"status"`
	Total       float64 `json:"total"`
	ItemBorrowd []Borrowd
}

func (borrowapi *BorrowAPI) Create() map[string]interface{} {

	t := time.Now()
	f := t.AddDate(0, 0, +3)
	data := Borrow{StartDate: t.Format("2006-01-02"), EndDate: f.Format("2006-01-02"), UserId: borrowapi.UserId, Status: "Belum Kembali", Total: borrowapi.Total}
	GetDB().Table("borrows").Create(&data)

	resp := u.Message(true, "success")
	resp["borrowapi"] = borrowapi
	return resp
}
