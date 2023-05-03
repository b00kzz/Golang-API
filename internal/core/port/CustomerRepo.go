package port

type CustomerRepo interface {
	GetAll() ([]Customer, error)
	GetById(id int) (*Customer, error)
	Create(Customer) (*Customer, error)
	Update(int, Customer) error
	Delete(int) error
}

type Customer struct {
	ID          uint   `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	DateOfBirth string `gorm:"column:date_Of_birth"`
	City        string `gorm:"column:city"`
	ZipCode     string `gorm:"column:zipcode"`
	Status      int    `gorm:"column:status"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
}

func (c Customer) TableName() string {
	return "tbl_customers"
}
