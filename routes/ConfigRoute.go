package routes

import (
	"github.com/NamDuongkiwi/Lab2/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigBookRouter(router *fiber.Router) {
	//Return all books
	(*router).Get("/", controller.GetAllBook)

	(*router).Get("/:id", controller.GetBookById)

	(*router).Delete("/:id", controller.DeleteBookById)
	(*router).Patch("/:id", controller.UpdateBook)

	(*router).Post("", controller.CreateBook)
}

func ConfigReviewRouter(route *fiber.Router){
	(*route).Get("/", controller.GetAllReviews)
	(*route).Post("/", controller.CreateReview)
}