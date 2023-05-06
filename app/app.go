package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"ticket/goapi/logs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	ginEngine *gin.Engine
}

func NewApp() AppPort {
	return &App{ginEngine: gin.New()}
}

func (a *App) GinEngine() gin.Engine {
	return *a.ginEngine
}

func (a *App) GinCors() {
	ginCors := gin.Default()
	ginCors.Use(cors.New(cors.DefaultConfig()))
}

func (a *App) UseMiddleware(handlers ...AppHandlerFunc) {
	ginHanlders := toGinHandlers(handlers...)
	a.ginEngine.Use(ginHanlders...)
}

func (a *App) Start(addr string) {
	logs.Info("Service starting...")
	srv := &http.Server{
		Addr:    addr,
		Handler: a.ginEngine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Error(err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Info("Service is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logs.Error(err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logs.Debug("timeout of 5 seconds.")
	}
	logs.Info("Service shut down successfully")
}

func toGinHandlers(handlers ...AppHandlerFunc) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			for _, h := range handlers {
				h(ctx)
			}
		},
	}
}
