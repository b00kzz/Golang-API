package domain

import "ticket/goapi/internal/core/port"

type ReviewSvc interface {
	GetAllReview() ([]ReviewRespone, error)
	GetReview(int) (*ReviewRespone, error)
	AddReview(ReviewRequest) (*ReviewRespone, error)
	UpdateReview(int, ReviewRequest) error
	UpdateStatusRev(int, StatusRev) error
	DeleteReview(int) error
}

type ReviewRequest struct {
	TicketId    uint   `json:"ticketid"`
	UserId      uint   `json:"userid"`
	RevRank     string `json:"revrank"`
	RevComment  string `json:"revcomment"`
	RevImage    string `json:"revimage"`
	Status      bool   `json:"status"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type ReviewRespone struct {
	TicketId    uint        `json:"ticketid"`
	RevId       uint        `json:"revid"`
	UserId      uint        `json:"userid"`
	RevRank     string      `json:"revrank"`
	RevComment  string      `json:"revcomment"`
	RevImage    string      `json:"revimage"`
	Status      bool   `json:"status"`
	CreatedBy   string      `json:"createdby"`
	CreatedDate string      `json:"createddate"`
	UpdatedBy   string      `json:"updatedby"`
	UpdatedDate string      `json:"updateddate"`
	Users       []port.User `json:"users"`
}

type StatusRev struct {
	Status bool `json:"status"`
}