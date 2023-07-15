package handler

import (
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/service"

	"github.com/gin-gonic/gin"
)

// type senderHdl struct {
// 	svc domain.SenderSvc
// }

// func NewSenderHdl(svc domain.SenderSvc) senderHdl {
// 	return senderHdl{
// 		svc: svc,
// 	}
// }

func SendMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := domain.SenderEmail{}
		err := c.BindJSON(&req)
		if err != nil {
			c.Error(err)
		}
		err = service.NewSenderSvc().SendEmail(req)
		if err != nil {
			c.Error(err)
		}

	}
}