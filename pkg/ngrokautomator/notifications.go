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

func sendNewUrl(params Email, config SmtpConfig) error {
	m := mail.NewMessage()
	//TODO better mail template
	m.SetHeader("From", params.sender)
	m.SetHeader("To", params.recipients)
	m.SetHeader("Subject", "Home Server")
	m.SetBody("text/html", "    <p>Hello,</p>\n    <p>Here is the link you requested:</p>\n    <a href=\""+params.variables[0]+"\">Click here to access the URL ==> </a>\n "+params.variables[0]+"    <p>If you have any questions or need further assistance, please hesitate to contact me, here is my fake number 25311184.</p>\n    <p>Best regards,</p>\n    <p>Your Name</p>")

	d := mail.NewDialer(config.server, config.port, config.username, config.password)

	if err := d.DialAndSend(m); err != nil {

		panic(err)
		return err
	}
	return nil
}

//TODO Warning mails
