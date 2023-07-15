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
func (h ticketHdl) GetAllByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("UserID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAllTicketID(id)
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
	if err != nil {
		c.Error(err)
	}
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
	if err != nil {
		c.Error(err)
	}
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

func (h ticketHdl) Search(c *gin.Context) {
	ticketname := c.Param("TicketName")

	res, _ := h.svc.Search(ticketname)
	c.JSON(http.StatusOK, res)
}

func (h *ticketHdl) UpdateStatusTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("TicketId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.StatusTicket{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateStatusTicket(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Ticket success!!",
	})

}
func (h *ticketHdl) UpdateSellStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("TicketId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.SellStatusTicket{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateSellStatus(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Ticket success!!",
	})

}
func (h *ticketHdl) UpdateCount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("TicketId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.TicketRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateCount(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Ticket success!!",
	})

}
