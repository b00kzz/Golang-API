package service

import (
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
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
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
		TicketId:    cust.TicketId,
		TicketName:  cust.TicketName,
		TicketType:  cust.TicketType,
		TicketPrice: cust.TicketPrice,
		TicketImage: cust.TicketImage,
		TicketDesc:  cust.TicketDesc,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}

func (r ticketSvc) AddTicket(req domain.TicketRequest) (*domain.TicketRespone, error) {
	currentTime := time.Now()
	cust := port.Ticket{
		TicketName:  req.TicketName,
		TicketType:  req.TicketType,
		TicketPrice: req.TicketPrice,
		TicketImage: req.TicketImage,
		TicketDesc:  req.TicketDesc,
		CreatedBy:   req.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save Ticket	")
	}
	resp := domain.TicketRespone{
		TicketName:  newCust.TicketName,
		TicketType:  newCust.TicketType,
		TicketPrice: newCust.TicketPrice,
		TicketImage: newCust.TicketImage,
		TicketDesc:  newCust.TicketDesc,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
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
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return &resp, nil
}
