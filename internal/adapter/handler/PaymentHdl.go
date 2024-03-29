package handler

import (
	"net/http"
	"strconv"
	"ticket/goapi/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type paymentHdl struct {
	svc domain.PaymentSvc
}

func NewPaymentHdl(svc domain.PaymentSvc) paymentHdl {
	return paymentHdl{
		svc: svc,
	}
}

func (h paymentHdl) SearchPayment(c *gin.Context) {
	name := c.Param("name")

	res, _ := h.svc.SearchPayment(name)
	c.JSON(http.StatusOK, res)
}

func (h paymentHdl) GetPayments(c *gin.Context) {
	res, err := h.svc.GetAllPayment()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h paymentHdl) GetPayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("PaymentId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetPayment(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h paymentHdl) AddPayment(c *gin.Context) {
	req := domain.PaymentRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	res, err := h.svc.AddPayment(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h paymentHdl) UpdatePayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("PaymentId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.PaymentRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdatePayment(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Payment success!!",
	})
}

func (h paymentHdl) DeletePayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("PaymentId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeletePayment(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Payment success!!",
	})
}
func (h paymentHdl) GetPaymentsId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("PaymentId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAllPaymentId(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h paymentHdl) GetPayUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("UserId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAllUserId(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
