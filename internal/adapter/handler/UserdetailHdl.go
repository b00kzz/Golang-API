package handler

import (
	"net/http"
	"strconv"
	"ticket/goapi/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type userDetailHdl struct {
	svc domain.UserDetailSvc
}

func NewUserDetailHdl(svc domain.UserDetailSvc) userDetailHdl {
	return userDetailHdl{
		svc: svc,
	}
}

func (h userDetailHdl) GetUserDetails(c *gin.Context) {
	res, err := h.svc.GetAllUserDetail()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h userDetailHdl) GetUserDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("UserDetailId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetUserDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h userDetailHdl) AddUserDetail(c *gin.Context) {
	req := domain.UserDetailRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	res, err := h.svc.AddUserDetail(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h userDetailHdl) UpdateUserDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("UserDetailId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.UserDetailRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateUserDetail(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update UserDetail success!!",
	})
}

func (h userDetailHdl) DeleteUserDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("UserDetailId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteUserDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete UserDetail success!!",
	})
}
