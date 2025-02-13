package main

import (
	"log"
	"os"

	"github.com/boPopov/tenant-api/api/database"
	"github.com/boPopov/tenant-api/api/routes"
	_ "github.com/boPopov/tenant-api/api/swaggerdocs"
	"github.com/gofiber/fiber/v2"
)

// @title Tenant API
// @version 1.0
// @description This is a sample API for managing tenants.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
func main() {

	database.ConnectDB()

	app := fiber.New() //Setting Fiber instance

	routes.SetRouths(app) //Setting rouths

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	log.Fatal(app.Listen(":" + port))
}
