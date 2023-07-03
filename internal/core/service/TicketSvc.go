package service

import (
	"fmt"
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"time"
)

type ticketSvc struct {
	repo port.TicketRepo
}

func NewTicketSvc(repo port.TicketRepo) domain.TicketSvc {
	return ticketSvc{
		repo: repo,
	}
}

func (s ticketSvc) GetAllTicket() ([]domain.TicketRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get ticket form DB")
	}
	resp := []domain.TicketRespone{}
	for _, c := range custs {
		resp = append(resp, domain.TicketRespone{
			TicketId:    c.TicketId,
			TicketName:  c.TicketName,
			TicketType:  c.TicketType,
			TicketPrice: c.TicketPrice,
			TicketImage: c.TicketImage,
			TicketDesc:  c.TicketDesc,
			Status:      c.Status,
			SellStatus:  c.SellStatus,
			Count:       c.Count,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
			TicketQr:    fmt.Sprintf("https://promptpay.io/0942710120/%d.png", c.TicketPrice),
		})

	}
	return resp, nil
}
func (s ticketSvc) GetAllTicketID(id int) ([]domain.TicketRespone, error) {
	custs, err := s.repo.GetAllByUserId(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get ticket form DB")
	}
	resp := []domain.TicketRespone{}
	for _, c := range custs {
		resp = append(resp, domain.TicketRespone{
			UserId:      c.UserId,
			TicketId:    c.TicketId,
			TicketName:  c.TicketName,
			TicketType:  c.TicketType,
			TicketPrice: c.TicketPrice,
			TicketImage: c.TicketImage,
			TicketDesc:  c.TicketDesc,
			Status:      c.Status,
			SellStatus:  c.SellStatus,
			Count:       c.Count,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
			TicketQr:    fmt.Sprintf("https://promptpay.io/0942710120/%d.png", c.TicketPrice),
		})

	}
	return resp, nil
}

func (s ticketSvc) GetTicket(id int) (*domain.TicketRespone, error) {

	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get Ticket form DB")
	}
	resp := domain.TicketRespone{
		UserId:      cust.UserId,
		TicketId:    cust.TicketId,
		TicketName:  cust.TicketName,
		TicketType:  cust.TicketType,
		TicketPrice: cust.TicketPrice,
		TicketImage: cust.TicketImage,
		TicketDesc:  cust.TicketDesc,
		TicketRepo:  cust.TicketRepo,
		Status:      cust.Status,
		Count:       cust.Count,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
		TicketQr:    fmt.Sprintf("https://promptpay.io/0942710120/%d.png", cust.TicketPrice),
	}
	return &resp, nil
}

func (r ticketSvc) AddTicket(req domain.TicketRequest) (*domain.TicketRespone, error) {
	currentTime := time.Now()
	cust := port.Ticket{
		UserId:      req.UserId,
		TicketName:  req.TicketName,
		TicketType:  req.TicketType,
		TicketPrice: req.TicketPrice,
		TicketImage: req.TicketImage,
		TicketDesc:  req.TicketDesc,
		TicketRepo:  req.TicketRepo,
		Status:      false,
		SellStatus:  false,
		Count:       0,
		CreatedBy:   req.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save Ticket	")
	}
	resp := domain.TicketRespone{
		UserId:      newCust.UserId,
		TicketName:  newCust.TicketName,
		TicketType:  newCust.TicketType,
		TicketPrice: newCust.TicketPrice,
		TicketImage: newCust.TicketImage,
		TicketDesc:  newCust.TicketDesc,
		Status:      newCust.Status,
		Count:       newCust.Count,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
		TicketQr:    fmt.Sprintf("https://promptpay.io/0942710120/%d.png", newCust.TicketPrice),
	}

	return &resp, nil
}

func (s ticketSvc) UpdateTicket(id int, req domain.TicketRequest) error {
	currentTime := time.Now()
	cust := port.Ticket{
		TicketName:  req.TicketName,
		TicketType:  req.TicketType,
		TicketPrice: req.TicketPrice,
		TicketImage: req.TicketImage,
		TicketDesc:  req.TicketDesc,
		UpdatedBy:   req.UpdatedBy,
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Ticket: ")
	}
	return nil
}
func (s ticketSvc) DeleteTicket(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete Ticket")
	}
	return nil
}

func (s ticketSvc) Search(ticketName string) (*[]domain.TicketRespone, error) {
	custs, err := s.repo.Search(ticketName)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get ticket form DB")
	}
	resp := []domain.TicketRespone{}
	for _, c := range custs {
		resp = append(resp, domain.TicketRespone{
			TicketId:    c.TicketId,
			TicketName:  c.TicketName,
			TicketType:  c.TicketType,
			TicketImage: c.TicketImage,
			TicketPrice: c.TicketPrice,
			TicketDesc:  c.TicketDesc,
			Status:      c.Status,
			SellStatus:  c.SellStatus,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return &resp, nil
}

func (s ticketSvc) UpdateStatusTicket(id int, req domain.StatusTicket) error {
	cust := port.Ticket{
		Status: req.Status,
	}
	err := s.repo.UpdateStatusTicket(id, cust.Status)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Status: ")
	}
	return nil
}
func (s ticketSvc) UpdateCount(id int, req domain.TicketRequest) error {
	currentTime := time.Now()
	cust := port.Ticket{
		Count:       +1,
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.UpdateCount(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Count: ")
	}
	return nil
}
func (s ticketSvc) UpdateSellStatus(id int, req domain.SellStatusTicket) error {
	cust := port.Ticket{
		SellStatus: req.SellStatus,
	}
	err := s.repo.UpdateSellStatus(id, cust.SellStatus)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update Status: ")
	}
	return nil
}
