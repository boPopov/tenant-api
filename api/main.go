package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New() //Setting Fiber instance
	//modules.SetRouths(app) //Setting rouths

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	log.Fatal(app.Listen(":" + port))
}
