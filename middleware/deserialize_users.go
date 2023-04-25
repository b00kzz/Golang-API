package middleware

import (
	"fmt"
	"ticket/goapi/errs"
	"ticket/goapi/infrastructure"

	// "ticket/goapi/infrastructure"
	"ticket/goapi/internal/core/port"
	"ticket/goapi/utils"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository port.RegisterRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := infrastructure.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error(), "sub": sub})
			return
		}

		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		errs.ErrorPanic(err_id)
		result, err := userRepository.GetById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Next()

	}
}
