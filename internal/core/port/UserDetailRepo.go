package port

type UserDetailRepo interface {
	GetAll() ([]UserDetail, error)
	GetById(id int) (*UserDetail, error)
	Create(UserDetail) (*UserDetail, error)
	Update(int, UserDetail) error
	Delete(int) error
}

type UserDetail struct {
	UserdeId     uint   `gorm:"primaryKey;autoIncrement"`
	UserId       uint   `gorm:"notnull"`
	RoleId       uint   `gorm:"notnull"`
	FirstName    string `gorm:"notnull"`
	LastName     string `gorm:"notnull"`
	Phone        string `gorm:"notnull;min=10;max=10"`
	Email        string `gorm:"notnull;min=10;max=10"`
	Avatar       string `gorm:"null;min=2"`
	RecordStatus string `gorm:"notnull"`
	CreatedBy    string `gorm:"notnull"`
	CreatedDate  string `gorm:"notnull"`
	UpdatedBy    string `gorm:"null"`
	UpdatedDate  string `gorm:"null"`
}

func (c UserDetail) TableName() string {
	return "tbl_userdetails"
}
