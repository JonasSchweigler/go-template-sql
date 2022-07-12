package routes

import "github.com/gofiber/fiber/v2"

type TestRoutes interface {
	InstallTestRoutes(app *fiber.App)
}
