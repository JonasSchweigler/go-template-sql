package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-template/repository"
)

//Add Handlers in here

type TestController interface {
	HelloWorld(ctx *fiber.Ctx) error
}

type testRepo struct {
	r repository.TestRepository
}

func NewTestRepository(repo repository.TestRepository) TestController {
	return &testRepo{r: repo}
}

//Routing Handlers for Controller

func (r *testRepo) HelloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World")
}
