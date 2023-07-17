package service

import (
	"fmt"
	"net/smtp"
	config "ticket/goapi/infrastructure"
	"ticket/goapi/internal/core/domain"
)

type sender struct{}

func NewSenderSvc() domain.SenderSvc {
	return sender{}
}

func (s sender) SendEmail(req domain.SenderEmail) error {
	auth := smtp.PlainAuth("", config.EnvSMTP_User(), config.EnvSMTP_Password(), config.EnvSMTP_Host())

	cust := domain.SenderEmail{
		Sender:  config.EnvSMTP_From(),
		SubJect: req.SubJect,
		Body:    req.Body,
		To:      req.To,
		Image:   req.Image,
	}

	body := fmt.Sprintf(`<p>จาก <b>Morlam Ticket</b>: ผลการซื้อ %s</p>`, cust.Body)
	body += fmt.Sprintf(`<img src="%s" width="100" height="100" />`, cust.Image)

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", cust.Sender)
	msg += fmt.Sprintf("To: %s\r\n", cust.To)
	msg += fmt.Sprintf("Subject: %s\r\n", cust.SubJect)
	msg += fmt.Sprintf("\r\n%s\r\n", body)

	err := smtp.SendMail(config.EnvSMTP_ADDR(), auth, config.EnvSMTP_From(), []string{cust.To}, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
