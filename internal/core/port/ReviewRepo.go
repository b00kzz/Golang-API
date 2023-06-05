package port

type ReviewRepo interface {
	GetAll() ([]Review, error)
	GetById(id int) (*Review, error)
	Create(Review) (*Review, error)
	Update(int, Review) error
	Delete(int) error
	UpdateStatusReview(int, bool) error
}

type Review struct {
	RevId       uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	TicketId    uint   `gorm:"notnull;type:int(10)"`
	UserId      uint   `gorm:"notnull;type:int(10)"`
	RevRank     string `gorm:"notnull;type:decimal(5)"`
	RevComment  string `gorm:"notnull;type:varchar(255)"`
	RevImage    string `gorm:"null;type:varchar(255)"`
	Status      bool   `gorm:"column:status;notnull;default:true"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
}

func (c Review) TableName() string {
	return "tbl_reviews"
}
