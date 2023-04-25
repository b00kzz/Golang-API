package domain

type CustomerSvc interface {
	GetAllCustomer() ([]CustomerResp, error)
	GetCustomer(int) (*CustomerResp, error)
	AddCustomer(CustomerReq) (*CustomerResp, error)
	UpdateCustomer(int, CustomerReq) error
	DeleteCustomer(int) error
}

type CustomerReq struct {
	Name        string `json:"name"`
	DateOfbirth string `json:"DateOfbirth"`
	City        string `json:"city"`
	ZipCode     string `json:"zipcode"`
	Status      int    `json:"status"`
}

type CustomerResp struct {
	CustomerId  uint   `json:"CustomerId"`
	Name        string `json:"name"`
	DateOfbirth string `json:"DateOfbirth"`
	City        string `json:"city"`
	ZipCode     string `json:"zipcode"`
	Status      int    `json:"status"`
}
