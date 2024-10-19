package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/BenMeredithConsult/locagri.worker.api/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri.worker.api/config"
)

type (
	SMSPayload struct {
		Message    string   `json:"message"`
		Recipients []string `json:"recipients"`
	}
	ArkeselPayload struct {
		Sender     string   `json:"sender"`
		Message    string   `json:"message"`
		Recipients []string `json:"recipients"`
	}
)

func newSMSService() gateways.Dispatcher {
	switch config.SMS().Gateway {
	case "arkesel":
		return newArkeselService()
	default:
		return nil
	}
}

type arkesel struct {
	APIKey string
	URL    string
	Sender string
}

func newArkeselService() gateways.Dispatcher {
	return &arkesel{
		APIKey: config.Arkesel().APIKey,
		URL:    config.Arkesel().URL,
		Sender: config.SMS().Sender,
	}
}

func (a *arkesel) Dispatch(payload []byte) {
	data := a.unmarshalMessage(payload)
	dataBody := ArkeselPayload{
		Sender:     a.Sender,
		Message:    data.Message,
		Recipients: a.formatRecipients(data.Recipients),
	}
	payloadBytes, err := json.Marshal(dataBody)
	if err != nil {
		log.Panicln(err)
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", a.URL, body)
	if err != nil {
		log.Panicln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", a.APIKey)
	fmt.Printf("Sending :%s: to %v\n", data.Message, dataBody.Recipients)
	//resp, resError := http.DefaultClient.Do(req)
	// if resError != nil {
	// 	log.Panicln(resError)
	// }
	// resp.Body.Close()
}

func (a *arkesel) formatRecipients(recipients []string) []string {
	var reps []string
	for _, recipient := range recipients {
		if strings.HasPrefix(recipient, "0") && len(recipient) == 10 {
			reps = append(reps, "233"+strings.Join(strings.Split(recipient, "")[1:], ""))
			continue
		}
		reps = append(reps, strings.Join(strings.Split(recipient, "")[1:], ""))
	}
	return reps
}
func (a *arkesel) unmarshalMessage(data []byte) *SMSPayload {
	var msg = new(SMSPayload)
	err := json.Unmarshal(data, msg)
	if err != nil {
		log.Panicln(err)
	}
	return msg
}
