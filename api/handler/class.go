package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rithikjain/motorq-task-backend/api/view"
	"github.com/rithikjain/motorq-task-backend/pkg/class"
)

func GetAllCourses(svc class.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		courses, err := svc.GetAllCourses()
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "All Courses Fetched",
			"courses": courses,
		})
	}
}

func MakeClassHandler(app *fiber.App, svc class.Service) {
	classGroup := app.Group("/api/class")
	classGroup.Get("/courses", GetAllCourses(svc))
}
