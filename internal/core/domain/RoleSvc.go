package domain

type RoleSvc interface {
	GetAllRole() ([]RoleRespone, error)
	GetRole(int) (*RoleRespone, error)
	AddRole(RoleRequest) (*RoleRespone, error)
	UpdateRole(int, RoleRequest) error
	DeleteRole(int) error
}

type RoleRequest struct {
	RoleName    string `json:"rolename" binding:"required"`
	RoleDesc    string `json:"roledesc" binding:"required"`
	Status      string `json:"status" binding:"required"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type RoleRespone struct {
	RoleId      uint   `json:"roleid" binding:"required"`
	RoleName    string `json:"rolename" binding:"required"`
	RoleDesc    string `json:"roledesc" binding:"required"`
	Status      string `json:"status" binding:"required"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}
