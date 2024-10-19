package notification

import (
	"bytes"
	"encoding/json"

	"github.com/BenMeredithConsult/locagri.worker.api/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri.worker.api/app/database"
)

type notification struct {
	db *database.Adapter
}

func NewNotification(db *database.Adapter) gateways.NotificationService {
	return &notification{
		db: db,
	}
}

func (n *notification) Mail() gateways.Dispatcher {
	return newMailer()
}
func (n *notification) SMS() gateways.Dispatcher {
	return newSMSService()
}
func (n *notification) DB() gateways.Dispatcher {
	return nil
}

func convertToBytes(data any) ([]byte, error) {
	dataBytes := new(bytes.Buffer)
	err := json.NewEncoder(dataBytes).Encode(data)
	if err != nil {
		return nil, err
	}
	return dataBytes.Bytes(), nil
}
