package domain

type ReviewSvc interface {
	GetAllReview() ([]ReviewRespone, error)
	GetReview(int) (*ReviewRespone, error)
	AddReview(ReviewRequest) (*ReviewRespone, error)
	UpdateReview(int, ReviewRequest) error
	DeleteReview(int) error
}

type ReviewRequest struct {
	RevRank     string `json:"revrank"`
	RevComment  string `json:"revcomment"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}

type ReviewRespone struct {
	RevId       uint   `json:"revid"`
	UserId      uint   `json:"userid"`
	RevRank     string `json:"revrank"`
	RevComment  string `json:"revcomment"`
	CreatedBy   string `json:"createdby"`
	CreatedDate string `json:"createddate"`
	UpdatedBy   string `json:"updatedby"`
	UpdatedDate string `json:"updateddate"`
}
