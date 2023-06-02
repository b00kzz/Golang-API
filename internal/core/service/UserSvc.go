package service

import (
	"errors"
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/infrastructure"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"ticket/goapi/utils"
	"time"
)

type registerSvc struct {
	repo port.RegisterRepo
}

func NewRegisterSvc(repo port.RegisterRepo) domain.RegisterSvc {
	return registerSvc{
		repo: repo,
	}
}

func (s registerSvc) GetAllUser() ([]domain.RegisterResp, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get user form DB")
	}
	resp := []domain.RegisterResp{}
	for _, c := range custs {
		resp = append(resp, domain.RegisterResp{
			ID:          c.ID,
			RoleId:      c.RoleId,
			Username:    c.Username,
			Password:    c.Password,
			Nickname:    c.Nickname,
			Email:       c.Email,
			Status:      c.Status,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s registerSvc) GetUser(id int) (*domain.RegisterResp, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get user form DB")
	}
	resp := domain.RegisterResp{
		ID:          cust.ID,
		RoleId:      cust.RoleId,
		Username:    cust.Username,
		// Password:    cust.Password,
		Nickname:    cust.Nickname,
		Email:       cust.Email,
		Avatar:      cust.Avatar,
		Status:      cust.Status,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedBy,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}

func (r registerSvc) AddUser(req domain.RegisterReq) (*domain.RegisterResp, error) {
	currentTime := time.Now()
	hashpwd, _ := utils.HashPassword(req.Password)
	cust := port.User{
		RoleId:      "User",
		Username:    req.Username,
		Password:    hashpwd,
		Nickname:    req.Nickname,
		Email:       req.Email,
		Status:      true,
		CreatedBy:   "System",
		CreatedDate: currentTime.Format(time.DateTime),
	}

	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save user	")
	}
	resp := domain.RegisterResp{
		RoleId:      newCust.RoleId,
		Username:    newCust.Username,
		Password:    newCust.Password,
		Nickname:    newCust.Nickname,
		Email:       newCust.Email,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s registerSvc) DeleteUser(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete user")
	}
	return nil
}

func (a registerSvc) Login(users domain.LoginReq) (string, error) {
	// Find username in database
	new_users, users_err := a.repo.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := infrastructure.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.ID, config.TokenSecret)
	errs.ErrorPanic(err_token)
	return token, nil

}

func (s registerSvc) GetProfile(username string) (*domain.RegisterResp, error) {
	cust, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get Profile form DB")
	}
	resp := domain.RegisterResp{
		ID:       cust.ID,
		RoleId:   cust.RoleId,
		Username: cust.Username,
		Password: cust.Password,
		Nickname: cust.Nickname,
		Email:    cust.Email,
		Avatar:   cust.Avatar,
		Status:   cust.Status,
	}
	return &resp, nil
}

func (s registerSvc) UpdateRole(id int, req domain.Role) error {
	cust := port.User{
		RoleId: req.RoleId,
	}
	err := s.repo.UpdateRole(id, cust.RoleId)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Role: ")
	}
	return nil
}

func (s registerSvc) UpdateStatus(id int, req domain.Status) error {
	cust := port.User{
		Status: req.Status,
	}
	err := s.repo.UpdateStatus(id, cust.Status)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Status: ")
	}
	return nil
}

func (s registerSvc) UpdateUser(id int, req domain.RegisterReq) error {
	currentTime := time.Now()
	hashpwd, _ := utils.HashPassword(req.Password)
	cust := port.User{
		Password:    hashpwd,
		Nickname:    req.Nickname,
		Email:       req.Email,
		Avatar:      req.Avatar,
		UpdatedBy:   "System",
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update user: ")
	}
	return nil
}
