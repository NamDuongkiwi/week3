package controller

import (
	_ "fmt"
	repo"github.com/NamDuongkiwi/Lab2/repository"
	"github.com/NamDuongkiwi/Lab2/model"
	"github.com/gofiber/fiber/v2"
)



func GetAllReviews(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}
func CreateReview(c *fiber.Ctx) error {
	var review model.Review
	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	book, err := repo.Books.FindBookById(review.BookId)
	if err != nil || book.Id ==0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "book not found",
			"book_id" : review.BookId,
		})
	}
	reviewId := repo.Reviews.CreateNewReview(&review)
	return c.JSON(fiber.Map{
		"message": "new review created",
		"reivew_id": reviewId,
	})
}