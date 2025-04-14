package mailer

import (
	"bytes"
	"embed"
	"github.com/materials-resources/store-api/app"
	mail "gopkg.in/gomail.v2"
	"html/template"
	"time"
)

//go:embed "templates"
var templateFS embed.FS

type Mailer struct {
	dialer *mail.Dialer
	sender string
}

// New initializes a new Mailer
func New(a *app.App) Mailer {

	dialer := mail.NewDialer(a.Config.Mailer.Host, a.Config.Mailer.Port, a.Config.Mailer.Username, a.Config.Mailer.Password)

	return Mailer{
		dialer: dialer,
		sender: a.Config.Mailer.Sender,
	}
}

// Send sends a new email to the provided recipient with the given template
func (m Mailer) Send(recipient, replyTo, templateFile string, data any) error {

	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	// set subject template
	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	// set plain body template
	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}

	// set html body template
	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	// create message
	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetHeader("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	// send message up to 3 times
	for i := 1; i <= 3; i++ {
		err = m.dialer.DialAndSend(msg)
		if nil == err {
			return nil
		}

		time.Sleep(500 * time.Millisecond)
	}

	return err

}
