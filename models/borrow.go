package models

import "time"

type Borrow struct {
	Id        uint      `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	UserId    uint      `json:"user_id"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"`
}

func GetBorrow(id string) *Borrow {
	borrow := &Borrow{}
	err := GetDB().Table("borrows").Where("id=?", id).First(borrow).Error
	if err != nil {
		return nil
	}
	return borrow
}
