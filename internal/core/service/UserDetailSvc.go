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
			FirstName:    c.FirstName,
			LastName:     c.LastName,
			Phone:        c.Phone,
			BankName:     c.BankName,
			BankId:       c.BankId,
			PersonCard:   c.PersonCard,
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
		FirstName:    cust.FirstName,
		LastName:     cust.LastName,
		Phone:        cust.Phone,
		BankName:     cust.BankName,
		BankId:       cust.BankId,
		PersonCard:   cust.PersonCard,
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
		UserId:       req.UserId,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		BankName:     req.BankName,
		BankId:       req.BankId,
		RecordStatus: "รออนุมัติ",
		PersonCard:   req.PersonCard,
		CreatedDate:  currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save UserDetail	")
	}
	resp := domain.UserDetailRespone{
		UserId:      newCust.UserId,
		FirstName:   newCust.FirstName,
		LastName:    newCust.LastName,
		Phone:       newCust.Phone,
		BankName:    newCust.BankName,
		BankId:      newCust.BankId,
		PersonCard:  newCust.PersonCard,
		CreatedDate: currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s userDetailSvc) UpdateUserDetail(id int, req domain.UserDetailRequest) error {
	currentTime := time.Now()
	cust := port.UserDetail{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        req.Phone,
		BankName:     req.BankName,
		BankId:       req.BankId,
		// PersonCard:   req.PersonCard,
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
