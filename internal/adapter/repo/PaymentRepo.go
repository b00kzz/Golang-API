package repo

import (
	"errors"
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) port.PaymentRepo {
	return paymentRepo{
		db: db,
	}
}

func (c paymentRepo) Search(name string) ([]port.Payment, error) {
	payments := []port.Payment{}
	result := c.db.Find(&payments, "ticket_name LIKE ? OR ticket_price LIKE ? OR ticket_desc LIKE ?", "%"+name+"%", "%"+name+"%", "%"+name+"%")

	if result.Error  != nil {
		return payments, errors.New("payment not found")
	}
	return payments, nil
}
func (c paymentRepo) GetAll() ([]port.Payment, error) {
	payments := []port.Payment{}
	err := c.db.Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (c paymentRepo) GetById(id int) (*port.Payment, error) {
	payment := port.Payment{}
	err := c.db.First(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (c paymentRepo) Create(payment port.Payment) (*port.Payment, error) {
	err := c.db.Create(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (c paymentRepo) Update(id int, payment port.Payment) error {
	err := c.db.Model(&port.Payment{}).Where("pay_id = ?", id).Updates(payment).Error
	if err != nil {
		return err
	}
	return nil
}

func (c paymentRepo) Delete(id int) error {
	err := c.db.Delete(&port.Payment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c paymentRepo) GetAllId(id int) ([]port.Payment, error) {
	payments := []port.Payment{}
	err := c.db.Where("by_id = ?", id).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (c paymentRepo) GetAllUserId(id int) ([]port.Payment, error) {
	payments := []port.Payment{}
	err := c.db.Where("user_id = ?", id).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}
