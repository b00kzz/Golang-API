package repo

import (
	"errors"
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type registerRepo struct {
	db *gorm.DB
}

func NewRegisterRepo(db *gorm.DB) port.RegisterRepo {
	return registerRepo{
		db: db,
	}
}

func (c registerRepo) Search(name string) ([]port.User, error) {
	user := []port.User{}
	result := c.db.Find(&user, "username LIKE ? OR nickname LIKE ? OR Email LIKE ?", "%"+name+"%", "%"+name+"%", "%"+name+"%")
	if result.Error != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

// มีการเปลี่ยนแปลง
func (c registerRepo) Create(userExist port.User) (*port.User, error) {
	_ = c.db.Where("username = ?", userExist.Username).First(&userExist).Error

	if userExist.ID > 0 {
		err := errors.New("username already exist")
		return nil, err
	} else {
		err := c.db.Create(&userExist).Error
		if err != nil {
			return &userExist, err
		}
	}
	return &userExist, nil
}

func (c registerRepo) Delete(id int) error {
	err := c.db.Delete(&port.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c registerRepo) GetAll() ([]port.User, error) {
	user := []port.User{}
	err := c.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c registerRepo) GetById(id int) (*port.User, error) {
	userID := port.User{}
	err := c.db.First(&userID, id).Error
	if err != nil {
		return nil, err
	}
	return &userID, nil
}

func (u registerRepo) FindByUsername(username string) (port.User, error) {
	var users port.User
	result := u.db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}

func (c registerRepo) UpdateRole(id int, role string) error {
	err := c.db.Model(&port.User{}).Where("user_id = ?", id).Update("role_id", role).Error
	if err != nil {
		return err
	}
	return nil
}
func (c registerRepo) UpdatePassword(id int, password string) error {
	err := c.db.Model(&port.User{}).Where("user_id = ?", id).Update("password", password).Error
	if err != nil {
		return err
	}
	return nil
}

func (c registerRepo) UpdateStatus(id int, status bool) error {
	err := c.db.Model(&port.User{}).Where("user_id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (c registerRepo) Update(id int, user port.User) error {
	err := c.db.Model(&port.User{}).Where("user_id = ?", id).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}
