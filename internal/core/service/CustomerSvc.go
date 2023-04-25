package service

import (
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"net/http"
)

type customerSvc struct {
	repo port.CustomerRepo
}

func NewCustomerSvc(repo port.CustomerRepo) domain.CustomerSvc {
	return customerSvc{
		repo: repo,
	}
}

func (s customerSvc) GetAllCustomer() ([]domain.CustomerResp, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get customer form DB")
	}
	resp := []domain.CustomerResp{}
	for _, c := range custs {
		resp = append(resp, domain.CustomerResp{
			CustomerId:  c.ID,
			Name:        c.Name,
			DateOfbirth: c.DateOfBirth,
			City:        c.City,
			ZipCode:     c.ZipCode,
		})

	}
	return resp, nil
}

func (s customerSvc) GetCustomer(id int) (*domain.CustomerResp, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get customer form DB")
	}
	resp := domain.CustomerResp{
		CustomerId:  cust.ID,
		Name:        cust.Name,
		DateOfbirth: cust.DateOfBirth,
		City:        cust.City,
		ZipCode:     cust.ZipCode,
	}
	return &resp, nil
}
func (s customerSvc) AddCustomer(req domain.CustomerReq) (*domain.CustomerResp, error) {
	cust := port.Customer{
		Name:        req.Name,
		DateOfBirth: req.DateOfbirth,
		City:        req.City,
		ZipCode:     req.ZipCode,
		Status:      req.Status,
	}
	newCust, err := s.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save customer")
	}
	resp := domain.CustomerResp{
		CustomerId:  newCust.ID,
		Name:        newCust.Name,
		DateOfbirth: newCust.DateOfBirth,
		City:        newCust.City,
		ZipCode:     newCust.ZipCode,
	}

	return &resp, nil
}
func (s customerSvc) UpdateCustomer(id int, req domain.CustomerReq) error {
	cust := port.Customer{
		Name:        req.Name,
		DateOfBirth: req.DateOfbirth,
		City:        req.City,
		ZipCode:     req.ZipCode,
		Status:      req.Status,
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update customer")
	}
	return nil
}
func (s customerSvc) DeleteCustomer(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete customer")
	}
	return nil
}
