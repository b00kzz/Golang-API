package port

type TicketRepo interface {
	GetAll() ([]Ticket, error)
	GetAllByUserId(int) ([]Ticket, error)
	GetById(id int) (*Ticket, error)
	Create(Ticket) (*Ticket, error)
	Update(int, Ticket) error
	Delete(int) error
	Search(string) ([]Ticket, error)
	UpdateStatusTicket(int, bool) error
	UpdateSellStatus(int, bool) error
}

type Ticket struct {
	TicketId    uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId      uint   `gorm:"notnull;type:int(10)"`
	TicketName  string `gorm:"notnull;type:varchar(150)"`
	TicketType  string `gorm:"notnull;type:varchar(50)"`
	TicketPrice string `gorm:"notnull;type:varchar(50)"`
	TicketImage string
	TicketDesc  string `gorm:"notnull;type:varchar(500)"`
	TicketRepo  string `gorm:"notnull"`
	Status      bool   `gorm:"column:status;notnull"`
	SellStatus  bool   `gorm:"notnull"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
}

func (c Ticket) TableName() string {
	return "tbl_Tickets"
}
