package service

import (
	"fmt"
	"net/smtp"
	"ticket/goapi/internal/core/domain"
)

type sender struct{}

func NewSenderSvc() domain.SenderSvc {
	return sender{}
}

func (s sender) SendEmail(req domain.SenderEmail) error {
	cust := domain.SenderEmail{
		SubJect: req.SubJect,
		Body:    req.Body,
		To:      req.To,
		Image:   req.Image,
	}
	auth := smtp.PlainAuth("", "qwasq110@gmail.com", "gqisfvfbozgqvvll", "smtp.gmail.com")
	msg := "Subject: " + cust.SubJect + "\n" + cust.Body
	// msg := "Subject: " + cust.SubJect + "\n" + cust.Body + "\n\n"
	// msg += "<img src=\"" + cust.Image + "\">"
	err := smtp.SendMail("smtp.gmail.com:587", auth, "qwasq110@gmail.com", []string{cust.To}, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
