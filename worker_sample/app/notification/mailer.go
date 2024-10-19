package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"github.com/BenMeredithConsult/locagri.worker.api/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri.worker.api/config"
	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type (
	MailerMessage struct {
		From        string
		FromName    string
		To          string
		Subject     string
		Attachments []string
		Data        any
		DataMap     map[string]any
		Template    string
	}
	mailer struct {
		SMTP        *mail.SMTPServer
		Mailer      string
		FromAddress string
		FromName    string
	}
)

func newMailer() gateways.Dispatcher {
	return &mailer{
		SMTP:        config.SMTPServer(),
		Mailer:      config.Mailer().Mailer,
		FromAddress: config.Mailer().FromAddress,
		FromName:    config.Mailer().FromName,
	}
}

func (m *mailer) Dispatch(payload []byte) {
	msg := m.unmarshalMessage(payload)
	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}

	msg.DataMap = data
	// build html mail
	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		log.Panicln("build:html:mail ", err)
	}

	// build plain text mail
	plainMessage, err := m.buildPlainTextMessage(msg)
	if err != nil {
		log.Panicln("build:text:mail ", err)
	}

	smtpClient, err := m.SMTP.Connect()
	if err != nil {
		log.Panicln("smtp:connection ", err)
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)
		}
	}
	err = email.Send(smtpClient)
	if err != nil {
		log.Panicln("send:mail ", err)
	}
}

func (m *mailer) unmarshalMessage(data []byte) *MailerMessage {
	var msg = new(MailerMessage)
	err := json.Unmarshal(data, msg)
	if err != nil {
		log.Panicln(err)
	}
	return msg
}

func (m *mailer) buildHTMLMessage(msg *MailerMessage) (string, error) {
	templateToRender := fmt.Sprintf("./app/template/emails/%s.html", msg.Template)

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *mailer) buildPlainTextMessage(msg *MailerMessage) (string, error) {
	templateToRender := fmt.Sprintf("./app/template/emails/%s.plain", msg.Template)

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()

	return plainMessage, nil
}

func (m *mailer) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}
