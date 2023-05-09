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

func (m *mdws) CORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		fmt.Println("OPTIONS")
	}
}
