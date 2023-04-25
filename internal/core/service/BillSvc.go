package service

import (
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"time"
)

type billSvc struct {
	repo port.BillRepo
}

func NewBillSvc(repo port.BillRepo) domain.BillSvc {
	return billSvc{
		repo: repo,
	}
}

func (s billSvc) GetAllBill() ([]domain.BillRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get users detail form DB")
	}
	resp := []domain.BillRespone{}
	for _, c := range custs {
		resp = append(resp, domain.BillRespone{
			BillID:      c.BillId,
			BillImg:     c.BillImg,
			BillStatus:  c.BillStatus,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s billSvc) GetBill(id int) (*domain.BillRespone, error) {

	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get Bill form DB")
	}
	resp := domain.BillRespone{
		BillID:      cust.BillId,
		BillImg:     cust.BillImg,
		BillStatus:  cust.BillStatus,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}

func (r billSvc) AddBill(req domain.BillRequest) (*domain.BillRespone, error) {
	currentTime := time.Now()
	cust := port.Bill{
		BillImg:     req.BillImg,
		BillStatus:  req.BillStatus,
		CreatedBy:   req.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save Bill	")
	}
	resp := domain.BillRespone{
		BillImg:     newCust.BillImg,
		BillStatus:  newCust.BillStatus,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s billSvc) UpdateBill(id int, req domain.BillRequest) error {
	currentTime := time.Now()
	cust := port.Bill{
		BillImg:     req.BillImg,
		BillStatus:  req.BillStatus,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Bill: ")
	}
	return nil
}
func (s billSvc) DeleteBill(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete Bill")
	}
	return nil
}
