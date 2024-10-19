package gateways

import "github.com/BenMeredithConsult/locagri.worker.api/app/database"

type (
	EventConsumer interface {
		Topics(topics []string) EventConsumer
		DB(db *database.Adapter) EventConsumer
		Listen()
	}
	Dispatcher interface {
		Dispatch(payload []byte)
	}
	NotificationService interface {
		SMS() Dispatcher
		Mail() Dispatcher
		DB() Dispatcher
	}
)
