package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	ErrorHandler(c *gin.Context)
	Logger(c *gin.Context)
}
