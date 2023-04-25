package domain

type UserDetailSvc interface {
	GetAllUserDetail() ([]UserDetailRespone, error)
	GetUserDetail(int) (*UserDetailRespone, error)
	AddUserDetail(UserDetailRequest) (*UserDetailRespone, error)
	UpdateUserDetail(int, UserDetailRequest) error
	DeleteUserDetail(int) error
}

type UserDetailRequest struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdby"`
	CreatedDate  string `json:"createddate"`
	UpdatedBy    string `json:"updatedby"`
	UpdatedDate  string `json:"updateddate"`
}

type UserDetailRespone struct {
	UserdeId     uint   `json:"userdeid"`
	UserId       uint   `json:"userid"`
	RoleId       uint   `json:"roleid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdby"`
	CreatedDate  string `json:"createddate"`
	UpdatedBy    string `json:"updatedby"`
	UpdatedDate  string `json:"updateddate"`
}
