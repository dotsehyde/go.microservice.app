package producer

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
)

const (
	TopicEmailNotification = "notification:email"
	TopicSMSNotification   = "notification:sms"
	TopicDBNotification    = "notification:db"
)

type producer struct {
	conn     *amqp.Connection
	exchange string
}

func NewProducer(conn *amqp.Connection, exchange string) gateways.EventProducer {
	prod := producer{
		conn:     conn,
		exchange: fmt.Sprintf("%s_topic", exchange),
	}
	channel, err := prod.conn.Channel()
	failOnError(err, "producer could not connect to RabbitMQ channel")
	err = declareExchange(channel, prod.exchange)
	failOnError(err, "producer declare exchange")
	return &prod
}

func (c *producer) Queue(queue string, payload any) {
	go func() {
		ch, err := c.conn.Channel()
		if err != nil {
			log.Panicln("error:channel:connection %w", err)
		}
		defer ch.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		payloadBytes, err := convertToBytes(payload)
		if err != nil {
			log.Panicln("error:cast:payload %w", err)
		}

		err = ch.PublishWithContext(
			ctx,
			c.exchange,
			queue,
			false,
			false,
			amqp.Publishing{
				ContentType:  "text/plain",
				DeliveryMode: amqp.Persistent,
				Body:         payloadBytes,
			},
		)
		if err != nil {
			log.Panicln("error:publish %w", err)
		}
	}()
}
