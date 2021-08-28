package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rithikjain/motorq-task-backend/api/view"
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"github.com/rithikjain/motorq-task-backend/pkg/student"
)

func AddStudent(svc student.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		studentBody := &models.Student{}
		if err := c.BodyParser(studentBody); err != nil {
			return view.Wrap(err, c)
		}

		err := svc.CreateStudent(studentBody)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Student Upserted",
		})
	}
}

func GetStudent(svc student.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		rollNo := c.Params("rollNo")

		stu, err := svc.GetStudentByRollNo(rollNo)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Student Fetched",
			"student": stu,
		})
	}
}

func MakeStudentHandler(app *fiber.App, svc student.Service) {
	studentGroup := app.Group("/api/student")
	studentGroup.Post("", AddStudent(svc))
	studentGroup.Get("/:rollNo", GetStudent(svc))
}
