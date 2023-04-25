package repo

import (
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type billRepo struct {
	db *gorm.DB
}

func NewBillRepo(db *gorm.DB) port.BillRepo {
	return billRepo{
		db: db,
	}
}

func (c billRepo) GetAll() ([]port.Bill, error) {
	bills := []port.Bill{}
	err := c.db.Find(&bills).Error
	if err != nil {
		return nil, err
	}
	return bills, nil
}

func (c billRepo) GetById(id int) (*port.Bill, error) {
	bill := port.Bill{}
	err := c.db.First(&bill, id).Error
	if err != nil {
		return nil, err
	}
	return &bill, nil
}

func (c billRepo) Create(bill port.Bill) (*port.Bill, error) {
	err := c.db.Create(&bill).Error
	if err != nil {
		return nil, err
	}
	return &bill, nil
}

func (c billRepo) Update(id int, bill port.Bill) error {
	err := c.db.Model(&port.Bill{}).Where("	bill_id = ?", id).Updates(bill).Error
	if err != nil {
		return err
	}
	return nil
}

func (c billRepo) Delete(id int) error {
	err := c.db.Delete(&port.Bill{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
