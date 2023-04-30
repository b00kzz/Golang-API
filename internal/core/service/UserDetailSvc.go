package service

import (
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"time"
)

type userDetailSvc struct {
	repo port.UserDetailRepo
}

func NewUserDetailSvc(repo port.UserDetailRepo) domain.UserDetailSvc {
	return userDetailSvc{
		repo: repo,
	}
}

func (s userDetailSvc) GetAllUserDetail() ([]domain.UserDetailRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get users detail form DB")
	}
	resp := []domain.UserDetailRespone{}
	for _, c := range custs {
		resp = append(resp, domain.UserDetailRespone{
			UserdeId:     c.UserdeId,
			UserId:       c.UserId,
			RoleId:       c.RoleId,
			FirstName:    c.FirstName,
			LastName:     c.LastName,
			Phone:        c.Phone,
			Email:        c.Email,
			Avatar:       c.Avatar,
			RecordStatus: c.RecordStatus,
			CreatedBy:    c.CreatedBy,
			CreatedDate:  c.CreatedDate,
			UpdatedBy:    c.UpdatedBy,
			UpdatedDate:  c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s userDetailSvc) GetUserDetail(id int) (*domain.UserDetailRespone, error) {

	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get UserDetail form DB")
	}
	resp := domain.UserDetailRespone{
		UserdeId:     cust.UserdeId,
		UserId:       cust.UserId,
		RoleId:       cust.RoleId,
		FirstName:    cust.FirstName,
		LastName:     cust.LastName,
		Phone:        cust.Phone,
		Email:        cust.Email,
		Avatar:       cust.Avatar,
		RecordStatus: cust.RecordStatus,
		CreatedBy:    cust.CreatedBy,
		CreatedDate:  cust.CreatedDate,
		UpdatedBy:    cust.UpdatedBy,
		UpdatedDate:  cust.UpdatedDate,
	}
	return &resp, nil
}

func (r userDetailSvc) AddUserDetail(req domain.UserDetailRequest) (*domain.UserDetailRespone, error) {
	currentTime := time.Now()
	cust := port.UserDetail{
		// UserId:       req.UserId,
		// RoleId:       req.RoleId,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		Email:        req.Email,
		Avatar:       req.Avatar,
		RecordStatus: req.RecordStatus,
		CreatedBy:    req.CreatedBy,
		CreatedDate:  currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save UserDetail	")
	}
	resp := domain.UserDetailRespone{
		// UserId:       newCust.UserId,
		FirstName:    newCust.FirstName,
		LastName:     newCust.LastName,
		Phone:        newCust.Phone,
		Email:        newCust.Email,
		Avatar:       newCust.Avatar,
		RecordStatus: newCust.RecordStatus,
		CreatedBy:    newCust.CreatedBy,
		CreatedDate:  currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s userDetailSvc) UpdateUserDetail(id int, req domain.UserDetailRequest) error {
	currentTime := time.Now()
	cust := port.UserDetail{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		Email:        req.Email,
		Avatar:       req.Avatar,
		RecordStatus: req.RecordStatus,
		UpdatedBy:    req.UpdatedBy,
		UpdatedDate:  currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update UserDetail: ")
	}
	return nil
}
func (s userDetailSvc) DeleteUserDetail(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete UserDetail")
	}
	return nil
}
