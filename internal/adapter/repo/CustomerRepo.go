package repo

import (
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) port.CustomerRepo {
	return customerRepo{
		db: db,
	}
}

func (c customerRepo) GetAll() ([]port.Customer, error) {
	customers := []port.Customer{}
	err := c.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c customerRepo) GetById(id int) (*port.Customer, error) {
	customer := port.Customer{}
	err := c.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c customerRepo) Create(customer port.Customer) (*port.Customer, error) {
	err := c.db.Create(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c customerRepo) Update(id int, customer port.Customer) error {
	err := c.db.Model(&port.Customer{}).Where("id = ?", id).Updates(customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (c customerRepo) Delete(id int) error {
	err := c.db.Delete(&port.Customer{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
