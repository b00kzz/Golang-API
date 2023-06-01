package handler

import (
	"net/http"
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/internal/core/service"

	"github.com/gin-gonic/gin"
)

func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		//upload
		formfile, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaResponse{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}
		uploadUrl, err := service.NewMediaUpload().FileUpload(domain.File{File: formfile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaResponse{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			domain.MediaResponse{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}

func RemoteUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url domain.Url

		//validate the request body
		if err := c.BindJSON(&url); err != nil {
			c.JSON(
				http.StatusBadRequest,
				domain.MediaResponse{
					StatusCode: http.StatusBadRequest,
					Message:    "error",
					Data:       map[string]interface{}{"data": err.Error()},
				})
			return
		}

		uploadUrl, err := service.NewMediaUpload().RemoteUpload(url)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaResponse{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			domain.MediaResponse{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
