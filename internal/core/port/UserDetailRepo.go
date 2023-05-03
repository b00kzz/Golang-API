package port

type UserDetailRepo interface {
	GetAll() ([]UserDetail, error)
	GetById(id int) (*UserDetail, error)
	Create(UserDetail) (*UserDetail, error)
	Update(int, UserDetail) error
	Delete(int) error
}

type UserDetail struct {
	UserdeId     uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	UserId       uint   `gorm:"notnull;type:int(10)"`
	RoleId       uint   `gorm:"notnull;type:int(10)"`
	FirstName    string `gorm:"notnull;type:varchar(100)"`
	LastName     string `gorm:"notnull;type:varchar(100)"`
	Phone        string `gorm:"notnull;type:varchar(10)"`
	Email        string `gorm:"notnull;type:varchar(100)"`
	Avatar       string `gorm:"null;type:varchar(100)"`
	RecordStatus string `gorm:"notnull;type:varchar(10)"`
	CreatedBy    string `gorm:"notnull;type:varchar(10)"`
	CreatedDate  string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy    string `gorm:"null;type:varchar(10)"`
	UpdatedDate  string `gorm:"null;type:varchar(20)"`
}

func (c UserDetail) TableName() string {
	return "tbl_userdetails"
}
