package handler

import (
	"net/http"
	"strconv"
	"ticket/goapi/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type billHdl struct {
	svc domain.BillSvc
}

func NewBillHdl(svc domain.BillSvc) billHdl {
	return billHdl{
		svc: svc,
	}
}

func (h billHdl) GetBills(c *gin.Context) {
	res, err := h.svc.GetAllBill()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h billHdl) GetBill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("BillId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetBill(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h billHdl) AddBill(c *gin.Context) {
	req := domain.BillRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	res, err := h.svc.AddBill(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h billHdl) UpdateBill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("BillId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.BillRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateBill(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Bill success!!",
	})
}

func (h billHdl) DeleteBill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("BillId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteBill(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Bill success!!",
	})
}
