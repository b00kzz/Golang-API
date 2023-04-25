package service

import (
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"time"
)

type roleSvc struct {
	repo port.RoleRepo
}

func NewRoleSvc(repo port.RoleRepo) domain.RoleSvc {
	return roleSvc{
		repo: repo,
	}
}

func (s roleSvc) GetAllRole() ([]domain.RoleRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get roles form DB")
	}
	resp := []domain.RoleRespone{}
	for _, c := range custs {
		resp = append(resp, domain.RoleRespone{
			RoleId:      c.RoleId,
			RoleName:    c.RoleName,
			RoleDesc:    c.RoleDesc,
			Status:      c.Status,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s roleSvc) GetRole(id int) (*domain.RoleRespone, error) {

	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get role form DB")
	}
	resp := domain.RoleRespone{
		RoleName:    cust.RoleName,
		RoleDesc:    cust.RoleDesc,
		Status:      cust.Status,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}

func (r roleSvc) AddRole(req domain.RoleRequest) (*domain.RoleRespone, error) {
	currentTime := time.Now()
	cust := port.Role{
		RoleName:    req.RoleName,
		RoleDesc:    req.RoleDesc,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save Role	")
	}
	resp := domain.RoleRespone{
		RoleName:    newCust.RoleName,
		RoleDesc:    newCust.RoleDesc,
		Status:      newCust.Status,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: newCust.CreatedDate,
	}

	return &resp, nil
}

func (s roleSvc) UpdateRole(id int, req domain.RoleRequest) error {
	currentTime := time.Now()
	cust := port.Role{
		RoleName:    req.RoleName,
		RoleDesc:    req.RoleDesc,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update user: ")
	}
	return nil
}
func (s roleSvc) DeleteRole(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete Role")
	}
	return nil
}
