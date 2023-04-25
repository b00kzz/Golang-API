package domain

type RegisterSvc interface {
	GetAllUser() ([]RegisterResp, error)
	GetUser(int) (*RegisterResp, error)
	AddUser(RegisterReq) (*RegisterResp, error)
	UpdateUser(int, RegisterReq) error
	DeleteUser(int) error
	Login(users LoginReq) (string, error)
}

type RegisterReq struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

type RegisterResp struct {
	ID       uint   `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

type LoginReq struct {
	Username string `validate:"required,max=200,min=2" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
