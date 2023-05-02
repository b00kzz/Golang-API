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
			Username:    c.Username,
			Password:    c.Password,
			Fullname:    c.Fullname,
			Email:       c.Email,
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
		Username:    cust.Username,
		Password:    cust.Password,
		Fullname:    cust.Fullname,
		Email:       cust.Email,
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
		Username:    req.Username,
		Password:    hashpwd,
		Fullname:    req.Fullname,
		Email:       req.Email,
		CreatedBy:   "User",
		CreatedDate: currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save user	")
	}
	resp := domain.RegisterResp{
		Username:    newCust.Username,
		Password:    newCust.Password,
		Fullname:    newCust.Fullname,
		Email:       newCust.Email,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s registerSvc) UpdateUser(id int, req domain.RegisterReq) error {
	currentTime := time.Now()
	hashpwd, _ := utils.HashPassword(req.Password)
	cust := port.User{
		Password:    hashpwd,
		Fullname:    req.Fullname,
		Email:       req.Email,
		UpdatedBy:   "User",
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update user: ")
	}
	return nil
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
		Username: cust.Username,
		Password: cust.Password,
		Fullname: cust.Fullname,
		Email:    cust.Email,
	}
	return &resp, nil
}
