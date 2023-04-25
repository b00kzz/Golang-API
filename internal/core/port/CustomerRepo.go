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
}

func (c Customer) TableName() string {
	return "tbl_customers"
}
