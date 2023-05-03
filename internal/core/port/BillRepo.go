package port

type BillRepo interface {
	GetAll() ([]Bill, error)
	GetById(id int) (*Bill, error)
	Create(Bill) (*Bill, error)
	Update(int, Bill) error
	Delete(int) error
}

type Bill struct {
	BillId      uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	BillImg     string `gorm:"notnull;type:varchar(100)"`
	BillStatus  string `gorm:"notnull;type:varchar(10)"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
}

func (c Bill) TableName() string {
	return "tbl_bills"
}
