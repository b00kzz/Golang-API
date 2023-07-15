package domain

type SenderSvc interface {
	SendEmail(SenderEmail) error
}

type SenderEmail struct {
	SubJect string `json:"subject"`
	Body    string `json:"body"`
	To      string `json:"to"`
	Image   string `json:"image"`
}
