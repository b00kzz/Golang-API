package service

import (
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"time"
)

type paymentSvc struct {
	repo port.PaymentRepo
}

func NewPaymentSvc(repo port.PaymentRepo) domain.PaymentSvc {
	return paymentSvc{
		repo: repo,
	}
}

func (s paymentSvc) GetAllPayment() ([]domain.PaymentRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get payment form DB")
	}
	resp := []domain.PaymentRespone{}
	for _, c := range custs {
		resp = append(resp, domain.PaymentRespone{
			PayId:       c.PayId,
			UserId:      c.UserId,
			TicketId:    c.TicketId,
			PaySlip:     c.PaySlip,
			PayStatus:   c.PayStatus,
			TicketName:  c.TicketName,
			TicketPrice: c.TicketPrice,
			TicketDesc:  c.TicketDesc,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s paymentSvc) GetPayment(id int) (*domain.PaymentRespone, error) {

	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get Payment form DB")
	}
	resp := domain.PaymentRespone{
		PayId:       cust.PayId,
		UserId:      cust.UserId,
		TicketId:    cust.TicketId,
		PaySlip:     cust.PaySlip,
		PayStatus:   cust.PayStatus,
		TicketName:  cust.TicketName,
		TicketPrice: cust.TicketPrice,
		TicketDesc:  cust.TicketDesc,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}

func (r paymentSvc) AddPayment(req domain.PaymentRequest) (*domain.PaymentRespone, error) {
	currentTime := time.Now()
	cust := port.Payment{
		UserId:      req.UserId,
		TicketId:    req.TicketId,
		PaySlip:     req.PaySlip,
		PayStatus:   "กำลังดำเนินการ",
		TicketName:  req.TicketName,
		TicketPrice: req.TicketPrice,
		TicketDesc:  req.TicketDesc,
		CreatedBy:   req.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save Payment	")
	}
	resp := domain.PaymentRespone{
		UserId:      newCust.UserId,
		TicketId:    newCust.TicketId,
		PaySlip:     newCust.PaySlip,
		PayStatus:   newCust.PayStatus,
		TicketName:  newCust.TicketName,
		TicketPrice: newCust.TicketPrice,
		TicketDesc:  newCust.TicketDesc,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s paymentSvc) UpdatePayment(id int, req domain.PaymentRequest) error {
	currentTime := time.Now()
	cust := port.Payment{
		PayStatus:   req.PayStatus,
		UpdatedBy:   "Admin",
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Payment: ")
	}
	return nil
}
func (s paymentSvc) DeletePayment(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete Payment")
	}
	return nil
}
