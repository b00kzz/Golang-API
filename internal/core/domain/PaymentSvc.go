package domain

type PaymentSvc interface {
	GetAllPayment() ([]PaymentRespone, error)
	GetPayment(int) (*PaymentRespone, error)
	AddPayment(PaymentRequest) (*PaymentRespone, error)
	UpdatePayment(int, PaymentRequest) error
	DeletePayment(int) error
}

type PaymentRequest struct {
	PayStatus   string `json:"paymentstatus"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type PaymentRespone struct {
	PayId       uint   `json:"payid"`
	UserId      uint   `json:"userid"`
	BillId      uint   `json:"billid"`
	TicketId    uint   `json:"ticketid"`
	PayStatus   string `json:"paymentstatus"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}