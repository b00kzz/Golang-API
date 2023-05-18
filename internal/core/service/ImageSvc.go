package service

import (
	"ticket/goapi/internal/core/domain"
	"ticket/goapi/utils"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file domain.File) (string, error)
	RemoteUpload(url domain.Url) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file domain.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := utils.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url domain.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := utils.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
