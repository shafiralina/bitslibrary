package models

import (
	u "bitslibrary/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type Borrow struct {
	gorm.Model
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	UserId    uint    `json:"user_id"`
	Status    string  `json:"status"`
	Total     float64 `json:"total"`
}

func (borrow *Borrow) Validate() (map[string]interface{}, bool) {

	if borrow.UserId == 0 {
		return u.Message(false, "UserId must be written"), false
	}
	return u.Message(true, "success"), true
}

func (borrow *Borrow) Create() map[string]interface{} {
	if resp, ok := borrow.Validate(); !ok {
		return resp
	}

	t := time.Now()
	f := t.AddDate(0, 0, +3)
	data := Borrow{StartDate: t.Format("2006-01-02"), EndDate: f.Format("2006-01-02"), UserId: borrow.UserId, Status: "Belum Kembali", Total: borrow.Total}
	GetDB().Table("borrows").Create(&data)

	resp := u.Message(true, "success")
	resp["borrow"] = borrow
	return resp
}
