package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3/middleware/recover"

	"github.com/BenMeredithConsult/locagri-apps/app/application/cache"
	"github.com/BenMeredithConsult/locagri-apps/app/application/producer"
	"github.com/BenMeredithConsult/locagri-apps/app/application/storage"
	"github.com/BenMeredithConsult/locagri-apps/app/framework/database"
	"github.com/BenMeredithConsult/locagri-apps/config"
	"github.com/BenMeredithConsult/locagri-apps/utils/env"
	"github.com/BenMeredithConsult/locagri-apps/utils/logger"
	"github.com/BenMeredithConsult/locagri-apps/utils/router"
	"github.com/gofiber/fiber/v3/middleware/cors"
	mlogger "github.com/gofiber/fiber/v3/middleware/logger"
)

func init() {
	env.Setup()
}

func App() {
	db := database.NewDB()
	// Run mirgation & Seed Languages
	database.MigrateDB(db)
	_ = database.Seeder(db)
	rdb := database.RedisDB()

	server := config.NewServer()

	storageSrv := storage.NewService(server.WG)
	cacheSrv := cache.NewService(rdb)
	mqConn, err := producer.Connect()

	if err != nil {
		logger.Log.Error("Failed to connect to RabbitMQ")
	}
	eventProducer := producer.NewProducer(mqConn, "worker")
	server.HTTP.Use(
		cors.New(
			cors.Config{
				AllowOrigins: "*",
			},
		),
	)

	server.HTTP.Use(recover.New())

	server.HTTP.Use(mlogger.New())

	// router.NewRouter(server.HTTP, db, dbConnections, rdb)
	router.NewRouter(server.HTTP, db, storageSrv, cacheSrv, rdb, eventProducer)

	go server.Run()
	go storageSrv.Listen()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(
		c, syscall.SIGINT, syscall.SIGTERM,
	) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	logger.Log.Info("Gracefully shutting down...")
	server.WG.Wait()
	storageSrv.Close()
	_ = server.HTTP.Shutdown()
	logger.Log.Info("Running cleanup tasks...")
	_ = db.DB.Close()
	logger.Log.Info("Application successful shutdown.")

}
