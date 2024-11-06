package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quiz.com/quiz/internal/entity"
	"quiz.com/quiz/internal/service"
)

type QuizController struct {
	quizService *service.QuizService
}

type GenerateQuizRequest struct {
	Name         string `json:"name"`
	Prompt       string `json:"prompt"`
	NumQuestions int    `json:"numQuestions"`
}

func Quiz(quizService *service.QuizService) QuizController {
	return QuizController{
		quizService: quizService,
	}
}

func (c QuizController) GetQuizById(ctx *fiber.Ctx) error {
	quizIdStr := ctx.Params("quizId")
	quizId, err := primitive.ObjectIDFromHex(quizIdStr)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	quiz, err := c.quizService.GetQuizById(quizId)
	if err != nil {
		return err
	}

	if quiz == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(quiz)
}

func (c *QuizController) GenerateAIQuiz(ctx *fiber.Ctx) error {
	fmt.Println("Received GenerateAIQuiz request")
	fmt.Printf("Raw body: %s\n", string(ctx.Body()))

	var body GenerateQuizRequest
	if err := ctx.BodyParser(&body); err != nil {
		fmt.Printf("Error parsing request body: %v\n", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid request body: %v", err),
		})
	}

	// Add validation for numQuestions
	if body.NumQuestions < 1 || body.NumQuestions > 20 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Number of questions must be between 1 and 20",
		})
	}

	fmt.Printf("Parsed request body: %+v\n", body)

	// Pass numQuestions to the service
	quiz, err := c.quizService.GenerateAIQuiz(body.Name, body.Prompt, body.NumQuestions)
	if err != nil {
		fmt.Printf("Error generating quiz: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(quiz)
}

type UpdateQuizRequest struct {
	Name      string                `json:"name"`
	Questions []entity.QuizQuestion `json:"questions"`
}

func (c *QuizController) CreateQuiz(ctx *fiber.Ctx) error {
	var body struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	quiz, err := c.quizService.CreateQuiz(body.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create quiz",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(quiz)
}

func (c QuizController) UpdateQuizById(ctx *fiber.Ctx) error {

	quizIdStr := ctx.Params("quizId")
	quizId, err := primitive.ObjectIDFromHex(quizIdStr)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var req UpdateQuizRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := c.quizService.UpdateQuiz(quizId, req.Name, req.Questions); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c QuizController) DeleteQuiz(ctx *fiber.Ctx) error {
	quizIdStr := ctx.Params("quizId")
	quizId, err := primitive.ObjectIDFromHex(quizIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid quiz ID format",
		})
	}

	err = c.quizService.DeleteQuiz(quizId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete quiz",
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := c.quizService.GetQuizzes()
	if err != nil {
		return err
	}

	return ctx.JSON(quizzes)
}
