package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/BenMeredithConsult/locagri-apps/app/application"
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
	SMSPayload struct {
		Message    string   `json:"message"`
		Recipients []string `json:"recipients"`
	}
)

func ComputeUniqueCode(delimiter string) (code string) {
	t := time.Now()
	st := strings.Split(fmt.Sprintf("%d", t.Year()), "")[2:]
	code = fmt.Sprintf("%s%s%s%sBR", application.OTP(4), st[0], st[1], delimiter)
	return
}

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
