package ngrokautomator

import (
	"github.com/go-mail/mail"
)

type Email struct {
	sender     string
	recipients string
	variables  []string
}

type SmtpConfig struct {
	server   string
	port     int
	username string
	password string
}

func sendNewUrl(params Email, cfg SmtpConfig) error {
	m := mail.NewMessage()
	m.SetHeader("From", params.sender)
	m.SetHeader("To", params.recipients)
	m.SetHeader("Subject", "New tunnel URL")
	m.SetBody("text/html", `<p>Hello,</p>
<p>A new tunnel has been established. Here is the URL:</p>
<p><a href="`+params.variables[0]+`">`+params.variables[0]+`</a></p>`)

	d := mail.NewDialer(cfg.server, cfg.port, cfg.username, cfg.password)
	return d.DialAndSend(m)
}
