package handler

import (
	"net/http"
	"strconv"
	"ticket/goapi/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type ticketHdl struct {
	svc domain.TicketSvc
}

func NewTicketHdl(svc domain.TicketSvc) ticketHdl {
	return ticketHdl{
		svc: svc,
	}
}

func (h ticketHdl) GetTickets(c *gin.Context) {
	res, err := h.svc.GetAllTicket()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h ticketHdl) GetTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("TicketId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetTicket(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h ticketHdl) AddTicket(c *gin.Context) {
	req := domain.TicketRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddTicket(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h ticketHdl) UpdateTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("TicketId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.TicketRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateTicket(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Ticket success!!",
	})
}

func (h ticketHdl) DeleteTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("TicketId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteTicket(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Ticket success!!",
	})
}
