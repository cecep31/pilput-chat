package main

import (
	"fmt"
	"log"
	apiroutes "pilput-chat/api/routes"
	"pilput-chat/config"
	"pilput-chat/pkg/book"
	"pilput-chat/ws/routes"

	"pilput-chat/ws/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, cancel, err := config.DatabaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())

	go handlers.RunHub()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome pilput-chat!"))
	})
	api := app.Group("/api")
	ws := app.Group("/ws")
	apiroutes.BookRouter(api, bookService)
	routes.ChatRouter(ws, bookService)
	defer cancel()
	log.Fatal(app.Listen(":8081"))
}
