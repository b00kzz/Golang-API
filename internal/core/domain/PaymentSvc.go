package domain

type PaymentSvc interface {
	GetAllPayment() ([]PaymentRespone, error)
	SearchPayment(string) (*[]PaymentRespone, error)
	GetAllPaymentId(int) ([]PaymentRespone, error)
	GetAllUserId(int) ([]PaymentRespone, error)
	GetPayment(int) (*PaymentRespone, error)
	AddPayment(PaymentRequest) (*PaymentRespone, error)
	UpdatePayment(int, PaymentRequest) error
	DeletePayment(int) error
}

type PaymentRequest struct {
	UserId      uint   `json:"userid"`
	TicketId    uint   `json:"ticketid"`
	ById        uint   `json:"byid"`
	PaySlip     string `json:"payslip"`
	PayStatus   string `json:"paymentstatus"`
	TicketName  string `json:"ticketname"`
	TicketPrice string `json:"ticketprice"`
	TicketDesc  string `json:"ticketdesc"`
	TicketRepo  string `json:"ticketrepo"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type PaymentRespone struct {
	PayId       uint   `json:"payid"`
	UserId      uint   `json:"userid"`
	PaySlip     string `json:"payslip"`
	TicketId    uint   `json:"ticketid"`
	PayStatus   string `json:"paymentstatus"`
	TicketName  string `json:"ticketname"`
	TicketPrice string `json:"ticketprice"`
	TicketDesc  string `json:"ticketdesc"`
	TicketRepo  string `json:"ticketrepo"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}
