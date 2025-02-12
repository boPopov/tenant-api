package main

import (
	"log"
	"os"

	"github.com/boPopov/tenant-api/api/database"
	"github.com/boPopov/tenant-api/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDB()

	app := fiber.New()    //Setting Fiber instance
	routes.SetRouths(app) //Setting rouths

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	log.Fatal(app.Listen(":" + port))
}
