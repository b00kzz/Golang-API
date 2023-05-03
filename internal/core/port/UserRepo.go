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
	ID          uint   `gorm:"column:user_id;type:int(10);primary_key;auto_increment"`
	Username    string `gorm:"column:username;type:varchar(50);notnull"`
	Password    string `gorm:"column:password;type:varchar(50);notnull"`
	Fullname    string `gorm:"column:fullname;type:varchar(200);notnull"`
	Email       string `gorm:"column:Email;type:varchar(100);notnull"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
}

func (u User) TableName() string {
	return "tbl_users"
}
