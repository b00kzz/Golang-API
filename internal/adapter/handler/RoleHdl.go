package handler

import (
	"net/http"
	"strconv"
	"ticket/goapi/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type roleHdl struct {
	svc domain.RoleSvc
}

func NewRoleHdl(svc domain.RoleSvc) roleHdl {
	return roleHdl{
		svc: svc,
	}
}

func (h roleHdl) GetRoles(c *gin.Context) {
	res, err := h.svc.GetAllRole()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h roleHdl) GetRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("RoleId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetRole(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h roleHdl) AddRole(c *gin.Context) {
	req := domain.RoleRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	res, err := h.svc.AddRole(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h roleHdl) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("RoleId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.RoleRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateRole(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Role success!!",
	})
}

func (h roleHdl) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("RoleId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteRole(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Role success!!",
	})
}
