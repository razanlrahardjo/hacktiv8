package routes

import (
	"github.com/razanlrahardjo/hacktiv8/app/controller"
	"github.com/razanlrahardjo/hacktiv8/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	services.InitDatabase()

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Get("/", controller.ApiIndexGet)

	// User Routing
	api.Post("/users", controller.PostUser)
	api.Get("/users", controller.GetUser)
	api.Put("/users/:id", controller.PutUser)
	api.Delete("/users/:id", controller.DeleteUser)

	// Todo Routing
	api.Post("/todos", controller.PostTodo)
	api.Get("/todos", controller.GetTodo)
	api.Get("/todos/:id", controller.GetTodoID)
	api.Put("/todos/:id", controller.PutTodo)
	api.Delete("/todos/:id", controller.DeleteTodo)

	// Status Routing
	api.Post("/status", controller.PostStatus)
	api.Get("/status", controller.GetStatus)
	api.Get("/status/:id", controller.GetStatusID)
	api.Put("/status/:id", controller.PutStatus)
	api.Delete("/status/:id", controller.DeleteStatus)

}
