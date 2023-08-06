package routes

import (
	"pilput-chat/pkg/book"
	"pilput-chat/ws/handlers"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func ChatRouter(app fiber.Router, service book.Service) {
	app.Get("/chat", handlers.Chat())
}
