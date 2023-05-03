package port

type RoleRepo interface {
	GetAll() ([]Role, error)
	GetById(id int) (*Role, error)
	Create(Role) (*Role, error)
	Update(int, Role) error
	Delete(int) error
}

type Role struct {
	RoleId      uint   `gorm:"primaryKey;autoIncrement;type:int(10)"`
	RoleName    string `gorm:"notnull;type:varchar(20)"`
	RoleDesc    string `gorm:"notnull;type:varchar(300)"`
	Status      string `gorm:"notnull;type:varchar(10)"`
	CreatedBy   string `gorm:"notnull;type:varchar(10)"`
	CreatedDate string `gorm:"notnull;type:varchar(20)"`
	UpdatedBy   string `gorm:"null;type:varchar(10)"`
	UpdatedDate string `gorm:"null;type:varchar(20)"`
}

func (c Role) TableName() string {
	return "tbl_roles"
}
