package port

type RegisterRepo interface {
	GetAll() ([]User, error)
	Search(string) ([]User, error)
	GetById(id int) (*User, error)
	Create(User) (*User, error)
	Update(int, User) error
	UpdateRole(int, string) error
	UpdateStatus(int, bool) error
	Delete(int) error
	FindByUsername(username string) (User, error)
}

type User struct {
	ID          uint   `gorm:"column:user_id;type:int(10);primary_key;auto_increment"`
	UserdeId    int    `gorm:"column:userde_id;type:int(10);notnull"`
	RoleId      string `gorm:"column:role_id;type:varchar(10);notnull"`
	Username    string `gorm:"column:username;type:varchar(50);notnull"`
	Password    string `gorm:"column:password;type:varchar(100);notnull"` //ถ้าเก็บขนาดน้อยไปจะไม่ได้
	Nickname    string `gorm:"column:nickname;type:varchar(200);notnull"`
	Email       string `gorm:"column:Email;type:varchar(100);notnull"`
	Status      bool   `gorm:"column:status;notnull;default:true"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
	Avatar      string `gorm:"null;type:longtext"`
}

func (u User) TableName() string {
	return "tbl_users"
}
