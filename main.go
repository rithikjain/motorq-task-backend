package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/rithikjain/motorq-task-backend/api/handler"
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"github.com/rithikjain/motorq-task-backend/pkg/student"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func dbConnect(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	// In the case of heroku
	if os.Getenv("onServer") == "True" {
		return gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	}
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)),
		&gorm.Config{},
	)

	return db, err
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("INFO: No PORT environment variable detected, defaulting to 4000")
		return "localhost:4000"
	}
	return ":" + port
}

func main() {
	if os.Getenv("onServer") != "True" {
		// Loading the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Setting up DB
	db, err := dbConnect(
		os.Getenv("dbHost"),
		os.Getenv("dbPort"),
		os.Getenv("dbUser"),
		os.Getenv("dbName"),
		os.Getenv("dbPass"),
		os.Getenv("sslmode"),
	)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	// Migrating tables
	err = db.AutoMigrate(
		&models.Student{},
		&models.Building{},
		&models.Course{},
		&models.Class{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New(fiber.Config{CaseSensitive: true})
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hey there looks like its working 🔥")
	})

	// Student
	studentRepo := student.NewRepo(db)
	studentSvc := student.NewService(studentRepo)
	handler.MakeStudentHandler(app, studentSvc)

	fmt.Println("Serving...")
	log.Fatal(app.Listen(GetPort()))
}
