package email

import (
	"bytes"
	"html/template"

	"github.com/wneessen/go-mail"
)

type Sender struct {
	client *mail.Client
	from   string
}

type VerificationEmailData struct {
	ConfirmURL string
}

func NewSender(
	smtpHost string,
	smtpPort int,
	smtpUser string,
	smtpPassword string,
	from string,
) (*Sender, error) {
	client, err := mail.NewClient(
		smtpHost,
		mail.WithPort(smtpPort),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(smtpUser),
		mail.WithPassword(smtpPassword),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return nil, err
	}

	return &Sender{
		client: client,
		from:   from,
	}, nil
}

func (s *Sender) SendVerificationEmail(
	to string,
	token string,
) error {
	tmpl, err := template.ParseFiles(
		"services/email-deliver/public/index.html",
	)
	if err != nil {
		return err
	}

	confirmURL := "http://localhost:5050/api/v1/users/login/verify?token=" + token

	data := VerificationEmailData{
		ConfirmURL: confirmURL,
	}

	var body bytes.Buffer

	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	msg := mail.NewMsg()

	if err := msg.From(s.from); err != nil {
		return err
	}

	if err := msg.To(to); err != nil {
		return err
	}

	msg.Subject("Подтверждение электронной почты")
	msg.SetBodyString(mail.TypeTextHTML, body.String())

	return s.client.DialAndSend(msg)
}
