package port

type RegisterRepo interface {
	GetAll() ([]User, error)
	GetById(id int) (*User, error)
	Create(User) (*User, error)
	Update(int, User) error
	Delete(int) error
	FindByUsername(username string) (User, error)
}

type User struct {
	ID          uint   `gorm:"column:user_id"`
	Username    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	Fullname    string `gorm:"column:fullname"`
	Email       string `gorm:"column:Email"`
	CreatedBy   string `gorm:"notnull"`
	CreatedDate string `gorm:"notnull"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (u User) TableName() string {
	return "tbl_users"
}
