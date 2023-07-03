package domain

type TicketSvc interface {
	GetAllTicket() ([]TicketRespone, error)
	GetAllTicketID(int) ([]TicketRespone, error)
	GetTicket(int) (*TicketRespone, error)
	AddTicket(TicketRequest) (*TicketRespone, error)
	UpdateTicket(int, TicketRequest) error
	DeleteTicket(int) error
	Search(string) (*[]TicketRespone, error)
	UpdateStatusTicket(int, StatusTicket) error
	UpdateSellStatus(int, SellStatusTicket) error
	UpdateCount(int, TicketRequest) error
}

type TicketRequest struct {
	UserId      uint   `json:"userid"`
	TicketName  string `json:"ticketname"`
	TicketType  string `json:"tickettype"`
	TicketPrice int `json:"ticketprice"`
	TicketImage string `json:"ticketimage"`
	TicketDesc  string `json:"ticketdesc"`
	TicketRepo  string `json:"ticketrepo"`
	Status      bool   `json:"status"`
	SellStatus  bool   `json:"sellstatus"`
	Count       int    `json:"count"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type TicketRespone struct {
	TicketId    uint   `json:"ticketid"`
	UserId      uint   `json:"userid"`
	TicketName  string `json:"ticketname"`
	TicketType  string `json:"tickettype"`
	TicketPrice int `json:"ticketprice"`
	TicketImage string `json:"ticketimage"`
	TicketDesc  string `json:"ticketdesc"`
	TicketRepo  string `json:"ticketrepo"`
	Status      bool   `json:"status"`
	SellStatus  bool   `json:"sellstatus"`
	Count       int    `json:"count"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
	TicketQr    string `json:"ticketqr"`
}

type StatusTicket struct {
	Status bool `json:"status"`
}
type SellStatusTicket struct {
	SellStatus bool `json:"sellstatus"`
}
