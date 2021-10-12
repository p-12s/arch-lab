package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

type MailTemplate struct {
	Id      int
	Subject string
	Body    string
}

func SendEmail(req common.NotificationSendReq) error {
	to := []string{
		req.Email,
	}

	mailTemplate := getMailTemplateById(req.TemplateId)
	request := Mail{
		Sender:  os.Getenv("EMAIL_NO_REPLY"),
		To:      to,
		Subject: mailTemplate.Subject,
		Body:    fmt.Sprintf(mailTemplate.Body, req.Context.FirstName, req.Context.LastName, req.Context.LoyaltyCardNumber),
	}

	message := BuildMessage(&request)
	smtpAuth := smtp.PlainAuth("", os.Getenv("SMTP_LOGIN"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))
	return smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), smtpAuth, os.Getenv("EMAIL_NO_REPLY"), to, []byte(message))
}

func BuildMessage(mail *Mail) string {
	var msg strings.Builder
	msg.WriteString("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n")
	msg.WriteString(fmt.Sprintf("From: %s\r\n", mail.Sender))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";")))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", mail.Subject))
	msg.WriteString(fmt.Sprintf("\r\n%s\r\n", mail.Body))

	return msg.String()
}

func getMailTemplateById(templateId int) MailTemplate {
	switch templateId {
	case 42:
		return MailTemplate{
			Id:      42,
			Subject: "Registration successfully completed",
			Body: `<div>
			<h1>Registration successfully completed!</h1>
			<p>Hello, %s %s!</p>
			<p>A discount card has been created for you,<br>here is its number: <b>%s</b></p>
			</div>`,
		}
	default:
		return MailTemplate{
			Id:      0,
			Subject: "This is our news for you",
			Body: `<div>
			<h1>Weekly news!</h1>
			<p>Look at <a href="https://site.ru">this</a>!</p>
			</div>`,
		}
	}
}
