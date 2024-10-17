package bootstrap

import (
	"os"
	"os/signal"
	"sample/config"
	"sample/utils/env"
	"sample/utils/logger"
	"sample/utils/router"
	"syscall"

	"github.com/gofiber/fiber/v2/middleware/cors"
	mlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	// os.Setenv("TZ", "Africa/Accra")
	env.Setup()
}

func App() {
	server := config.NewServer()

	server.HTTP.Use(
		cors.New(
			cors.Config{
				AllowOrigins: "*",
			},
		),
	)

	server.HTTP.Use(recover.New())

	server.HTTP.Use(mlogger.New())
	router.NewRouter(server.HTTP)
	go server.Run()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(
		c, syscall.SIGINT, syscall.SIGTERM,
	) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	logger.Log.Info("Gracefully shutting down...")
	server.WG.Wait()
	_ = server.HTTP.Shutdown()
	logger.Log.Info("Running cleanup tasks...")
	// for _, c := range dbConnections {
	// 	_ = c.DB.Close()
	// }
	logger.Log.Info("Application successful shutdown.")

}
