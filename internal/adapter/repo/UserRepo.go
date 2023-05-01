package repo

import (
	"errors"
	"ticket/goapi/internal/core/port"

	"golang.org/x/crypto/bcrypt"
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

func (r registerRepo) Create(user port.User) (*port.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	usernew := port.User{Username: user.Username, Password: string(hash), Fullname: user.Fullname, Email: user.Email, CreatedBy: user.CreatedBy, CreatedDate: user.CreatedDate, RoleID: user.RoleID}
	err := r.db.Create(&usernew).Error
	r.db.Model(&usernew).Update("userde_id", usernew.ID)
	if user.ID > 0 {
		return nil, err
	}
	return &usernew, nil
}

func (c registerRepo) Update(id int, user port.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	usernew := port.User{Username: user.Username, Password: string(hash), Fullname: user.Fullname, Email: user.Email, CreatedBy: user.CreatedBy, CreatedDate: user.CreatedDate, RoleID: user.RoleID}
	err := c.db.Model(&port.User{}).Where("user_id = ?", id).Updates(usernew).Error
	if err != nil {
		return err
	}
	return nil
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
