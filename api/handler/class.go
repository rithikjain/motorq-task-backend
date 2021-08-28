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

func GetAllClassesForACourse(svc class.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		courseID := c.Params("courseID")

		classes, err := svc.GetAllClassesForACourse(courseID)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Classes Fetched",
			"classes": classes,
		})
	}
}

func AddStudentToClass(svc class.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		studentID := c.Params("studentID")
		classID := c.Params("classID")

		err := svc.AddClassStudent(studentID, classID)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Student added to class",
		})
	}
}

func RemoveStudentFromClass(svc class.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		studentID := c.Params("studentID")
		classID := c.Params("classID")

		err := svc.RemoveClassStudent(studentID, classID)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Student removed from class",
		})
	}
}

func GetStudentTimetable(svc class.Service) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		studentID := c.Params("studentID")

		classes, err := svc.GetEnrolledClasses(studentID)
		if err != nil {
			return view.Wrap(err, c)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Enrolled classes fetched",
			"classes": classes,
		})
	}
}

func MakeClassHandler(app *fiber.App, svc class.Service) {
	classGroup := app.Group("/api/class")
	classGroup.Get("/courses", GetAllCourses(svc))
	classGroup.Get("/fetch/:courseID", GetAllClassesForACourse(svc))
	classGroup.Get("/map/:courseID", GetAllClassesForACourse(svc))
	classGroup.Post("/addStudent/:studentID/:classID", AddStudentToClass(svc))
	classGroup.Delete("/removeStudent/:studentID/:classID", RemoveStudentFromClass(svc))
	classGroup.Get("/timetable/:studentID", GetStudentTimetable(svc))
}
