package app

import "github.com/gin-gonic/gin"

type AppHandlerFunc func(*gin.Context)

type AppPort interface {
	Start(string)
	UseMiddleware(handlers ...AppHandlerFunc)
	GinEngine() gin.Engine
}
