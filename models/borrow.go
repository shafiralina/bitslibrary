package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Borrow struct {
	Id        uint      `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	UserId    uint      `json:"user_id"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"`
}

func (Borrow) TableName() string {
	return "borrows"
}

func GetBorrow(conn *gorm.DB, id string) *Borrow {
	borrow := &Borrow{}
	err := conn.Where("id=?", id).First(borrow).Error
	if err != nil {
		return nil
	}

	return borrow
}
