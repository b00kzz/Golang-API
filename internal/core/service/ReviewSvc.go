package service

import (
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/port"
	"time"
)

type reviewSvc struct {
	repo port.ReviewRepo
}

func NewReviewSvc(repo port.ReviewRepo) domain.ReviewSvc {
	return reviewSvc{
		repo: repo,
	}
}

func (s reviewSvc) GetAllReview() ([]domain.ReviewRespone, error) {
	custs, err := s.repo.GetAll()
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get review form DB")
	}
	resp := []domain.ReviewRespone{}
	for _, c := range custs {
		resp = append(resp, domain.ReviewRespone{
			RevId:       c.RevId,
			TicketId:    c.TicketId,
			UserId:      c.UserId,
			RevRank:     c.RevRank,
			RevComment:  c.RevComment,
			RevImage:    c.RevImage,
			Status:      c.Status,
			CreatedBy:   c.CreatedBy,
			CreatedDate: c.CreatedDate,
			UpdatedBy:   c.UpdatedBy,
			UpdatedDate: c.UpdatedDate,
		})

	}
	return resp, nil
}

func (s reviewSvc) GetReview(id int) (*domain.ReviewRespone, error) {
	cust, err := s.repo.GetById(id)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot get Review form DB")
	}
	resp := domain.ReviewRespone{
		TicketId:    cust.TicketId,
		UserId:      cust.UserId,
		RevRank:     cust.RevRank,
		RevComment:  cust.RevComment,
		CreatedBy:   cust.CreatedBy,
		CreatedDate: cust.CreatedDate,
		UpdatedBy:   cust.UpdatedBy,
		UpdatedDate: cust.UpdatedDate,
	}
	return &resp, nil
}

func (r reviewSvc) AddReview(req domain.ReviewRequest) (*domain.ReviewRespone, error) {
	currentTime := time.Now()
	cust := port.Review{
		TicketId:    req.TicketId,
		UserId:      req.UserId,
		RevRank:     req.RevRank,
		RevComment:  req.RevComment,
		RevImage:    req.RevImage,
		Status:      true,
		CreatedBy:   req.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}
	newCust, err := r.repo.Create(cust)
	if err != nil {
		return nil, errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot save Review	")
	}
	resp := domain.ReviewRespone{
		TicketId:    newCust.TicketId,
		UserId:      newCust.UserId,
		RevRank:     newCust.RevRank,
		RevComment:  newCust.RevComment,
		Status:      newCust.Status,
		CreatedBy:   newCust.CreatedBy,
		CreatedDate: currentTime.Format(time.DateTime),
	}

	return &resp, nil
}

func (s reviewSvc) UpdateReview(id int, req domain.ReviewRequest) error {
	currentTime := time.Now()
	cust := port.Review{
		RevRank:     req.RevRank,
		RevComment:  req.RevComment,
		UpdatedDate: currentTime.Format(time.DateTime),
	}
	err := s.repo.Update(id, cust)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update review")
	}
	return nil
}
func (s reviewSvc) UpdateStatusRev(id int, req domain.StatusRev) error {
	cust := port.Review{
		Status: req.Status,
	}
	err := s.repo.UpdateStatusReview(id, cust.Status)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot update review")
	}
	return nil
}
func (s reviewSvc) DeleteReview(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errs.New(http.StatusInternalServerError, "80001", errs.SystemErr, "Cannot delete Review")
	}
	return nil
}
