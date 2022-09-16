package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/angmeng/task_app/config"
	"github.com/angmeng/task_app/pkg/helpers"
	"github.com/angmeng/task_app/routes"
	"github.com/angmeng/task_app/stores"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./templates", ".html")
	config.Init()
	app := fiber.New(fiber.Config{
		AppName:     fmt.Sprintf("Task App %s", helpers.GetAppEnv()),
		Prefork:     helpers.IsProduction(),
		IdleTimeout: 5 * time.Second,
		Views:       engine, //set as render engine
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).SendString(err.Error())
		},
	})

	config.Setup(app)
	routes.DrawRoutes(app)
	sess := stores.ConnectPG()
	defer sess.Close()
	fmt.Println("database connected")

	go func() {
		port := config.Val.ServerPort
		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
