package domain

type RegisterSvc interface {
	GetAllUser() ([]RegisterResp, error)
	SearchUser(string) (*[]RegisterResp, error)
	GetUser(int) (*RegisterResp, error)
	AddUser(RegisterReq) (*RegisterResp, error)
	UpdateUser(int, RegisterReq) error
	DeleteUser(int) error
	Login(users LoginReq) (string, error)
	GetProfile(string) (*RegisterResp, error)
	UpdateStatus(int, Status) error
	UpdateRole(int, Role) error
}

type RegisterReq struct {
	RoleId      string `json:"roleid"`
	Username    string `validate:"required,min=2,max=100" json:"username"`
	Password    string `validate:"required,min=2,max=100" json:"password"`
	Nickname    string `json:"nickname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Avatar      string `json:"avatar"`
	Status      bool   `json:"status"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type RegisterResp struct {
	ID          uint   `json:"userid" binding:"required"`
	RoleId      string `json:"roleid"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	OldPasswd   string `json:"oldpassword" binding:"required"`
	Nickname    string `json:"nickname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Avatar      string `json:"avatar"`
	Status      bool   `json:"status"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type LoginReq struct {
	Username string `validate:"required,max=200,min=2" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type LoginResponse struct {
	TokenType string      `json:"token_type"`
	Token     string      `json:"token"`
	Bearer    interface{} `json:"bearer,omitempty"`
}

type Response struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	User        interface{} `json:"user,omitempty"`
	AccessToken string      `json:"accessToken,omitempty"`
}

type Role struct {
	RoleId string `json:"roleid"`
}
type Status struct {
	Status bool `json:"status"`
}

type Password struct {
	Password  string `json:"password"`
	OldPasswd string `json:"oldpassword"`
}
