package consumer

import (
	"fmt"
	"log"

	"github.com/BenMeredithConsult/locagri.worker.api/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri.worker.api/app/database"
	"github.com/BenMeredithConsult/locagri.worker.api/app/notification"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	conn     *amqp.Connection
	exchange string
	topics   []string
	notify   gateways.NotificationService
}

func New(conn *amqp.Connection, exchange string) gateways.EventConsumer {
	consumer := consumer{
		conn:     conn,
		exchange: fmt.Sprintf("%s_topic", exchange),
	}
	channel, err := consumer.conn.Channel()
	consumer.failOnError(err, "Consumer could not connect to RabbitMQ channel")
	err = consumer.declareExchange(channel, consumer.exchange)
	consumer.failOnError(err, "Consumer declare exchange")
	return &consumer
}

// DB implements gateways.EventConsumer.
func (c *consumer) DB(db *database.Adapter) gateways.EventConsumer {
	c.notify = notification.NewNotification(db)
	return c
}

func (c *consumer) Topics(topics []string) gateways.EventConsumer {
	c.topics = topics
	return c
}

func (c *consumer) Listen() {
	ch, err := c.conn.Channel()
	if err != nil {
		log.Panicln("error:event_consumer %w", err)
	}
	defer ch.Close()

	q, err := c.declareRandomQueue(ch)
	if err != nil {
		log.Panicln("error:event_consumer %w", err)
	}

	for _, topic := range c.topics {
		if err := ch.QueueBind(
			q.Name,
			topic,
			c.exchange,
			false,
			nil,
		); err != nil {
			log.Panicln("error:event_consumer %w", err)
		}
	}
	if err := ch.Qos(1, 0, false); err != nil {
		log.Panicln("error:event_consumer %w", err)
	}
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Panicln("error:event_consumer %w", err)
	}
	forever := make(chan bool)
	go func() {
		for message := range messages {
			go c.executeTask(message.RoutingKey, message.Body)
		}
	}()
	log.Printf("Waiting for message [Exchange, Queue] [%s, %s]\n", c.exchange, q.Name)
	<-forever
}
func (c *consumer) executeTask(routeKey string, payload []byte) {
	switch routeKey {
	case "notification:email":
		c.notify.Mail().Dispatch(payload)
	case "notification:sms":
		// fmt.Printf("Executing - SMS: %s\n", payload)
		c.notify.SMS().Dispatch(payload)
	case "notification:db":
		// c.notify.DB().Dispatch(payload)
	}
}

func (c *consumer) declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
}
func (c *consumer) declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
func (c *consumer) failOnError(err error, msg ...string) {
	if err != nil {
		log.Panicf("%s: %s", msg[0], err)
	}
}
