package repo

import (
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type userdetailRepo struct {
	db *gorm.DB
}

func NewUserDetailRepo(db *gorm.DB) port.UserDetailRepo {
	return userdetailRepo{
		db: db,
	}
}

func (c userdetailRepo) GetAll() ([]port.UserDetail, error) {
	userdetails := []port.UserDetail{}
	err := c.db.Find(&userdetails).Error
	if err != nil {
		return nil, err
	}
	return userdetails, nil
}

func (c userdetailRepo) GetById(id int) (*port.UserDetail, error) {
	userdetail := port.UserDetail{}
	err := c.db.First(&userdetail, id).Error
	if err != nil {
		return nil, err
	}
	return &userdetail, nil
}

func (c userdetailRepo) Create(userdetail port.UserDetail) (*port.UserDetail, error) {
	err := c.db.Create(&userdetail).Error
	if err != nil {
		return nil, err
	}
	return &userdetail, nil
}

func (c userdetailRepo) Update(id int, userdetail port.UserDetail) error {
	err := c.db.Model(&port.UserDetail{}).Where("userde_id = ?", id).Updates(userdetail).Error
	if err != nil {
		return err
	}
	return nil
}

func (c userdetailRepo) Delete(id int) error {
	err := c.db.Delete(&port.UserDetail{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
