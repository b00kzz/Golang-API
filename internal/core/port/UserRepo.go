package port

type RegisterRepo interface {
	GetAll() ([]User, error)
	GetById(id int) (*User, error)
	Create(User) (*User, error)
	Update(int, User) error
	Delete(int) error
	FindByUsername(username string) (User, error)
	// GetProfileById(id int) (*User, error)
}

type User struct {
	ID       uint   `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Fullname string `gorm:"column:fullname"`
	Avatar   string `gorm:"column:avatar"`
}

func (u User) TableName() string {
	return "tbl_users"
}
