package bootstrap

import (
	"fmt"
	"log"
	"log/slog"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BenMeredithConsult/locagri.worker.api/app/consumer"
	"github.com/BenMeredithConsult/locagri.worker.api/app/database"
	"github.com/BenMeredithConsult/locagri.worker.api/config"
	"github.com/BenMeredithConsult/locagri.worker.api/utils/env"
	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	env.Setup()
}
func App() {
	amqpConn, err := connect()
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Println("Listener connected to RabbitMQ...")
	db := database.NewDB()
	consume := consumer.New(amqpConn, "worker").
		Topics([]string{"notification:db", "notification:email", "notification:sms", "notification:all"}).
		DB(db)

	go consume.Listen()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(
		c, syscall.SIGINT, syscall.SIGTERM,
	) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	log.Println("Listener Application Gracefully shutting down...")
	_ = db.DB.Close()
	amqpConn.Close()
	log.Println("Listener Application successful shutdown.")
}

func connect() (*amqp.Connection, error) {
	var count int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s", config.App().RabbitMQ))
		if err != nil {
			log.Println("RabbitMQ not yet ready")
			slog.Error("RabbitMQ Err:" + err.Error())
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

func failOnError(err error, msg ...string) {
	if err != nil {
		log.Panicf("%s: %s", msg[0], err)
	}
}
