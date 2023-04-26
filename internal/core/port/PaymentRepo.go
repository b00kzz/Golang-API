package port

type PaymentRepo interface {
	GetAll() ([]Payment, error)
	GetById(id int) (*Payment, error)
	Create(Payment) (*Payment, error)
	Update(int, Payment) error
	Delete(int) error
}

type Payment struct {
	PayId       uint   `gorm:"primaryKey;autoIncrement"`
	UserId      uint   `gorm:"notnull"`
	BillId      uint   `gorm:"notnull"`
	TicketId    uint   `gorm:"notnull"`
	PayStatus   string `gorm:"notnull;default=กำลังดำเนินการ"`
	CreatedBy   string `gorm:"notnull"`
	CreatedDate string `gorm:"notnull"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Payment) TableName() string {
	return "tbl_payments"
}
