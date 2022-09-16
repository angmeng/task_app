package config

import (
	"io"
	"os"

	"github.com/angmeng/task_app/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

var Val *Config

type Config struct {
	DatabaseName     string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	FiberENV         string
	ServerPort       string
}

func Init() {
	AppRoot := helpers.AppRoot()
	// Load .env file
	if err := godotenv.Load(AppRoot + "/.env"); err != nil {
		panic("Error loading .env file")
	}

	Val = &Config{
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASS"),
		ServerPort:       os.Getenv("SERVER_PORT"),
	}
}

func Setup(app *fiber.App) {
	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${pid} ${latency} - ${status} ${method} ${path} | ${queryParams}\n${body}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Kuala_Lumpur",
		Output:     io.MultiWriter(os.Stdout),
	}))

	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/dont_compress"
		},
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Use(cors.New(cors.Config(cors.Config{
		Next:             func(c *fiber.Ctx) bool { return false },
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Access-Control-Allow-Origin, Host, X-Host, Referer",
		ExposeHeaders:    "",
		MaxAge:           0,
		AllowCredentials: true,
	})))
}
