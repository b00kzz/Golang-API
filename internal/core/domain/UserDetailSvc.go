package domain

type UserDetailSvc interface {
	GetAllUserDetail() ([]UserDetailRespone, error)
	GetUserDetail(int) (*UserDetailRespone, error)
	AddUserDetail(UserDetailRequest) (*UserDetailRespone, error)
	UpdateUserDetail(int, UserDetailRequest) error
	DeleteUserDetail(int) error
}

type UserDetailRequest struct {
	UserId       uint   `json:"userid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Phone        string `json:"phone"`
	BankName     string `json:"bankname"`
	BankId       string `json:"bankid"`
	PersonCard   string `json:"personcard"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdby"`
	CreatedDate  string `json:"createddate"`
	UpdatedBy    string `json:"updatedby"`
	UpdatedDate  string `json:"updateddate"`
}

type UserDetailRespone struct {
	UserdeId     uint   `json:"userdeid"`
	UserId       uint   `json:"userid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Phone        string `json:"phone"`
	BankName     string `json:"bankname"`
	BankId       string `json:"bankid"`
	PersonCard   string `json:"personcard"`
	RecordStatus string `json:"recordstatus"`
	CreatedBy    string `json:"createdby"`
	CreatedDate  string `json:"createddate"`
	UpdatedBy    string `json:"updatedby"`
	UpdatedDate  string `json:"updateddate"`
}
