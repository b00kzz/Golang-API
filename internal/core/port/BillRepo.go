package port

type BillRepo interface {
	GetAll() ([]Bill, error)
	GetById(id int) (*Bill, error)
	Create(Bill) (*Bill, error)
	Update(int, Bill) error
	Delete(int) error
}

type Bill struct {
	BillId      uint   `gorm:"primaryKey;autoIncrement"`
	BillImg     string `gorm:"notnull"`
	BillStatus  string `gorm:"notnull"`
	CreatedBy   string `gorm:"notnull"`
	CreatedDate string `gorm:"notnull"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Bill) TableName() string {
	return "tbl_bills"
}
