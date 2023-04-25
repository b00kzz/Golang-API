package domain

type BillSvc interface {
	GetAllBill() ([]BillRespone, error)
	GetBill(int) (*BillRespone, error)
	AddBill(BillRequest) (*BillRespone, error)
	UpdateBill(int, BillRequest) error
	DeleteBill(int) error
}

type BillRequest struct {
	BillImg     string `json:"billimg"`
	BillStatus  string `json:"billstatus"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type BillRespone struct {
	BillID      uint   `json:"billid"`
	BillImg     string `json:"billimg"`
	BillStatus  string `json:"billstatus"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}