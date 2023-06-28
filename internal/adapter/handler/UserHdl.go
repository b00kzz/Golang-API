package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"ticket/goapi/errs"
	"ticket/goapi/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type registerHdl struct {
	svc domain.RegisterSvc
}

func NewRegisterHdl(svc domain.RegisterSvc) registerHdl {
	return registerHdl{
		svc: svc,
	}
}

func (h *registerHdl) GetUsers(c *gin.Context) {
	// currentUser := c.MustGet("currentUser").(port.User)
	users, _ := h.svc.GetAllUser()
	webResponse := domain.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		User:    users,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h registerHdl) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetUser(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h registerHdl) AddUser(c *gin.Context) {
	req := domain.RegisterReq{}
	c.ShouldBindJSON(&req) //+มา
	res, err := h.svc.AddUser(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h registerHdl) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.RegisterReq{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateUser(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update User success!!",
	})
}

func (h registerHdl) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteUser(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User success!!",
	})
}

func (h *registerHdl) Login(ctx *gin.Context) {
	loginRequest := domain.LoginReq{}
	err := ctx.ShouldBindJSON(&loginRequest)
	errs.ErrorPanic(err)

	token, err_token := h.svc.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := domain.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	users, _ := h.svc.GetProfile(loginRequest.Username)
	resp := domain.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := domain.Response{
		Code:        200,
		Status:      "Ok",
		Message:     "Successfully log in!",
		AccessToken: resp.Token,
		User:        users,
	}
	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (h registerHdl) GetProfile(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(string)
	users, _ := h.svc.GetProfile(currentUser)
	webResponse := domain.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		User:    users,
	}
	c.JSON(http.StatusOK, webResponse)

}

func (h *registerHdl) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.Role{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateRole(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update User success!!",
	})

}
func (h *registerHdl) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.Status{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateStatus(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Status success!!",
	})

}
func (h *registerHdl) UpdatePasswd(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.Password{}
	err = c.BindJSON(&req)
	err = h.svc.UpdatePasswd(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Password success!!",
	})

}

func (h registerHdl) SearchUser(c *gin.Context) {
	name := c.Param("name")

	res, _ := h.svc.SearchUser(name)
	c.JSON(http.StatusOK, res)
}
