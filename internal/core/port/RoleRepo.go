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
	RoleName    string `gorm:"notnull"`
	RoleDesc    string `gorm:"notnull"`
	Status      string `gorm:"notnull"`
	CreatedBy   string `gorm:"notnull"`
	CreatedDate string `gorm:"notnull"`
	UpdatedBy   string `gorm:"null"`
	UpdatedDate string `gorm:"null"`
}

func (c Role) TableName() string {
	return "tbl_roles"
}
