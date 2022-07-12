package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-template/controllers"
)

type testRoutes struct {
	controller controllers.TestController
}

func NewTestRoutes(testController controllers.TestController) TestRoutes {
	return &testRoutes{testController}
}

func (r *testRoutes) InstallTestRoutes(app *fiber.App) {
	app.Get("/api/test", r.controller.HelloWorld)
}
