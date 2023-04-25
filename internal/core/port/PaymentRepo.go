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
	UserId      uint   `gorm:"notnull;minlength=2;maxlength=100"`
	BillId      uint   `gorm:"notnull;minlength=2;maxlength=100"`
	TicketId    uint   `gorm:"notnull;minlength=2;maxlength=100"`
	PayStatus   string `gorm:"notnull;default=กำลังดำเนินการ"`
	CreatedBy   string `gorm:"notnull;minlength=2;maxlength=100"`
	CreatedDate string `gorm:"notnull;minlength=2;maxlength=100"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Payment) TableName() string {
	return "tbl_payments"
}
