package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go-template/controllers"
	"go-template/db"
	"go-template/repository"
	"go-template/routes"
	"log"
)

func main() {

	//TODO: Please change the IP Name to ensure that the program won't crash

	conn := db.NewConnection("dbname")
	defer conn.Close()
	app := fiber.New()
	app.Use(logger.New())

	//Provides Statistics about the CPU and so on

	app.Get("/monitor", monitor.New(monitor.ConfigDefault))

	//Test Routers
	testRepo := repository.NewTestRepository(conn)
	testController := controllers.NewTestRepository(testRepo)
	testRoutes := routes.NewTestRoutes(testController)
	testRoutes.InstallTestRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
