package repo

import (
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) port.RoleRepo {
	return roleRepo{
		db: db,
	}
}

func (c roleRepo) GetAll() ([]port.Role, error) {
	roles := []port.Role{}
	err := c.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (c roleRepo) GetById(id int) (*port.Role, error) {
	role := port.Role{}
	err := c.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (c roleRepo) Create(role port.Role) (*port.Role, error) {
	err := c.db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (c roleRepo) Update(id int, role port.Role) error {
	err := c.db.Model(&port.Role{}).Where("role_id = ?", id).Updates(role).Error
	if err != nil {
		return err
	}
	return nil
}

func (c roleRepo) Delete(id int) error {
	err := c.db.Delete(&port.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
