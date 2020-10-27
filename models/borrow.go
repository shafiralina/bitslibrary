package models

import (
	"github.com/jinzhu/gorm"
)

type Borrow struct {
	gorm.Model
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	UserId    uint    `json:"user_id"`
	Status    string  `json:"status"`
	Total     float64 `json:"total"`
}

func GetBorrow(id string) *Borrow {
	borrow := &Borrow{}
	err := GetDB().Table("borrows").Where("id=?", id).First(borrow).Error
	if err != nil {
		return nil
	}
	return borrow
}
