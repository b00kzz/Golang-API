package domain

type TicketSvc interface {
	GetAllTicket() ([]TicketRespone, error)
	GetTicket(int) (*TicketRespone, error)
	AddTicket(TicketRequest) (*TicketRespone, error)
	UpdateTicket(int, TicketRequest) error
	DeleteTicket(int) error
}

type TicketRequest struct {
	TicketName  string `json:"ticketname"`
	TicketType  string `json:"tickettype"`
	TicketPrice string `json:"ticketprice"`
	TicketDesc  string `json:"ticketdesc"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type TicketRespone struct {
	TicketId    uint   `json:"ticketid"`
	TicketName  string `json:"ticketname"`
	TicketType  string `json:"tickettype"`
	TicketPrice string `json:"ticketprice"`
	TicketDesc  string `json:"ticketdesc"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}