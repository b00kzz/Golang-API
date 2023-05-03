package middleware

import (
	"fmt"
	"net/http"
	"ticket/goapi/errs"
	"ticket/goapi/logs"
	"time"

	"github.com/gin-gonic/gin"
)

type mdws struct {
}

func New() Middleware {
	return &mdws{}
}

func (m *mdws) ErrorHandler(c *gin.Context) {
	c.Next()
	errors := c.Errors
	for _, err := range errors {
		switch e := err.Err.(type) {
		case *errs.Errs:
			c.AbortWithStatusJSON(e.HTTPStatusCode, e)
		case error:
			c.AbortWithStatusJSON(http.StatusInternalServerError, e)
		}
		return
	}
}

func (m *mdws) Logger(c *gin.Context) {
	t := time.Now()
	logs.Info(fmt.Sprintf("BEGIN| %v | %s%s", c.Request.Method, c.Request.Host, c.Request.URL))
	logs.Info("Executing the proceed....")
	c.Next()
	latency := time.Since(t)
	status := c.Writer.Status()
	logs.Info(fmt.Sprintf("END | %v | %v | %v | %s%s", c.Request.Method, status, latency, c.Request.Host, c.Request.URL))
}
