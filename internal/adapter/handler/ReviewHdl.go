package handler

import (
	"net/http"
	"strconv"
	"ticket/goapi/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type reviewHdl struct {
	svc domain.ReviewSvc
}

func NewReviewHdl(svc domain.ReviewSvc) reviewHdl {
	return reviewHdl{
		svc: svc,
	}
}

func (h reviewHdl) GetReviews(c *gin.Context) {
	res, err := h.svc.GetAllReview()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h reviewHdl) GetReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ReviewId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetReview(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h reviewHdl) AddReview(c *gin.Context) {
	req := domain.ReviewRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddReview(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h reviewHdl) UpdateReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ReviewId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.ReviewRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateReview(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Review success!!",
	})
}

func (h reviewHdl) DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ReviewId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteReview(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Review success!!",
	})
}
