package models

type Borrowd struct {
	BorrowId uint    `json:"borrow_id"`
	BookId   int     `json:"book_id"`
	Qty      int     `json:"qty"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}

func (Borrowd) TableName() string {
	return "borrowds"
}

func GetDetailBorrow(id string) []*Borrowd {
	borrowd := make([]*Borrowd, 0)
	err := GetDB().Where("borrow_id=?", id).Find(&borrowd).Error
	if err != nil {
		return nil
	}
	defer db.Close()
	return borrowd
}
