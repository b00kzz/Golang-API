package port

type PaymentRepo interface {
	GetAll() ([]Payment, error)
	GetById(id int) (*Payment, error)
	Create(Payment) (*Payment, error)
	Update(int, Payment) error
	Delete(int) error
}

type Payment struct {
	PayId       uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId      uint   `gorm:"notnull;type:int(10)"`
	TicketId    uint   `gorm:"notnull;type:int(10)"`
	PaySlip     string `gorm:"notnull"`
	PayStatus   string `gorm:"notnull;default=กำลังดำเนินการ;type:varchar(20)"`
	TicketName  string `gorm:"notnull;type:varchar(150)"`
	TicketPrice string `gorm:"notnull;type:varchar(50)"`
	TicketDesc  string `gorm:"notnull;type:varchar(500)"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
}

func (c Payment) TableName() string {
	return "tbl_payments"
}
