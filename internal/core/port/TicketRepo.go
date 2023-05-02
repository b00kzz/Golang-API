package port

type TicketRepo interface {
	GetAll() ([]Ticket, error)
	GetById(id int) (*Ticket, error)
	Create(Ticket) (*Ticket, error)
	Update(int, Ticket) error
	Delete(int) error
	Search(string) ([]Ticket, error)
}

type Ticket struct {
	TicketId    uint   `gorm:"primaryKey;autoIncrement"`
	TicketName  string `gorm:"notnull"`
	TicketType  string `gorm:"notnull"`
	TicketPrice string `gorm:"notnull"`
	TicketDesc  string `gorm:"notnull"`
	CreatedBy   string `gorm:"notnull"`
	CreatedDate string `gorm:"notnull"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Ticket) TableName() string {
	return "tbl_Tickets"
}