package routes

import (
	"github.com/angmeng/task_app/internal/handlers/home"
	"github.com/angmeng/task_app/internal/handlers/tasks"
	"github.com/gofiber/fiber/v2"
)

func DrawRoutes(app *fiber.App) {
	app.Static("/public", "./public")
	app.Get("/", home.Index)

	api := app.Group("/api")
	api.Get("/tasks", tasks.Index)
	api.Post("/tasks", tasks.Create)
	api.Put("/tasks/:id", tasks.Update)
	api.Delete("/tasks/:id", tasks.Delete)
}
