package port

type RoleRepo interface {
	GetAll() ([]Role, error)
	GetById(id int) (*Role, error)
	Create(Role) (*Role, error)
	Update(int, Role) error
	Delete(int) error
}

type Role struct {
	RoleId      uint   `gorm:"primaryKey;autoIncrement"`
	RoleName    string `gorm:"notnull;minlength=2;maxlength=100"`
	RoleDesc    string `gorm:"notnull;minlength=2;maxlength=100"`
	Status      string `gorm:"notnull;minlength=2;maxlength=100"`
	CreatedBy   string `gorm:"notnull;minlength=2;maxlength=100"`
	CreatedDate string `gorm:"notnull;minlength=2;maxlength=100"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Role) TableName() string {
	return "tbl_roles"
}
