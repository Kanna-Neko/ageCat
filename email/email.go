package email

import (
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

func SendEmail(toEmailPath []string, fromEmailPath string, emailPassword string, smtpPath string, port string, data []byte) error{
	e := &email.Email{
		To:      toEmailPath,
		From:    "ageCat<" + fromEmailPath + ">",
		Subject: "ageCat notification",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    data,
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(smtpPath+":"+port, smtp.PlainAuth("", fromEmailPath, emailPassword, smtpPath))
	if err != nil {
		return err
	}
	return nil
}
