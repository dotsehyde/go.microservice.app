package producer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/BenMeredithConsult/locagri-apps/config"
)

type (
	WorkerPayload struct {
		Event string `json:"event"`
		Data  []byte `json:"data"`
	}
)

func Connect() (*amqp.Connection, error) {
	var count int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection
	for {
		conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s", config.App().RabbitMQ))
		if err != nil {
			log.Println("RabbitMQ not yet ready")
			count++
		} else {
			connection = conn
			break
		}
		if count > 5 {
			log.Println(err)
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		log.Println("RabbitMQ not yet ready")
		time.Sleep(backOff)
		continue
	}
	return connection, nil
}

func declareExchange(ch *amqp.Channel, exchange string) error {
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

func failOnError(err error, msg ...string) {
	if err != nil {
		log.Panicf("%s: %s", msg[0], err)
	}
}

func convertToBytes(data any) ([]byte, error) {
	dataBytes := new(bytes.Buffer)
	err := json.NewEncoder(dataBytes).Encode(data)
	if err != nil {
		return nil, err
	}
	return dataBytes.Bytes(), nil
}
