package models

import "github.com/jinzhu/gorm"

type Borrowd struct {
	gorm.Model
	BorrowId uint    `json:"borrow_id"`
	BookId   int     `json:"book_id"`
	Qty      int     `json:"qty"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}

func GetDetailBorrow(id string) []*Borrowd {
	borrowd := make([]*Borrowd, 0)
	err := GetDB().Table("borrowds").Where("borrow_id=?", id).Find(&borrowd).Error
	if err != nil {
		return nil
	}
	return borrowd
}
