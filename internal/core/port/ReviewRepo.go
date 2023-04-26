package port

type ReviewRepo interface {
	GetAll() ([]Review, error)
	GetById(id int) (*Review, error)
	Create(Review) (*Review, error)
	Update(int, Review) error
	Delete(int) error
}

type Review struct {
	RevId       uint   `gorm:"primaryKey;autoIncrement"`
	UserId      uint   `gorm:"notnull"`
	RevRank     string `gorm:"notnull;type:decimal(5)"`
	RevComment  string `gorm:"notnull;type:varchar(255)"`
	CreatedBy   string `gorm:"notnull"`
	CreatedDate string `gorm:"notnull"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Review) TableName() string {
	return "tbl_reviews"
}
