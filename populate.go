package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rithikjain/motorq-task-backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func connectDB(host, port, user, dbname, password, sslmode string, choice int) (*gorm.DB, error) {
	// In the case of Heroku
	if choice == 2 {
		return gorm.Open(postgres.Open(os.Getenv("HEROKU_DATABASE_URL")), &gorm.Config{})
	}

	// Default Case
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)),
		&gorm.Config{},
	)

	return db, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Which DB to migrate:\n 1) Local \n 2) Heroku")
	var choice int
	_, _ = fmt.Scanln(&choice)

	db, err := connectDB(
		os.Getenv("dbHost"),
		os.Getenv("dbPort"),
		os.Getenv("dbUser"),
		os.Getenv("dbName"),
		os.Getenv("dbPass"),
		os.Getenv("sslmode"),
		choice,
	)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	//var courses = []models.Course{
	//	{
	//		ID:   "CSE1006",
	//		Name: "Blockchain and Cryptocurrency Technologies",
	//	},
	//	{
	//		ID:   "CSE2010",
	//		Name: "Advanced C Programming",
	//	},
	//	{
	//		ID:   "CSE3045",
	//		Name: "Mathematical Modeling for Data Science",
	//	},
	//	{
	//		ID:   "CSE3054",
	//		Name: "Data Mining: Concepts and Techniques",
	//	},
	//}
	//db.Create(courses)

	//var buildings = []models.Building{
	//	{
	//		Name: "SJT",
	//		Location: models.Location{
	//			Lat: 12.971063,
	//			Lon: 79.163851,
	//		},
	//	},
	//	{
	//		Name: "TT",
	//		Location: models.Location{
	//			Lat: 12.9707336,
	//			Lon: 79.158584,
	//		},
	//	},
	//	{
	//		Name: "SMV",
	//		Location: models.Location{
	//			Lat: 12.9692483,
	//			Lon: 79.1568129,
	//		},
	//	},
	//	{
	//		Name: "MGR",
	//		Location: models.Location{
	//			Lat: 12.9689666,
	//			Lon: 79.1564537,
	//		},
	//	},
	//}
	//db.Create(buildings)

	var classes = []models.Class{
		{
			CourseID:     "CSE1006",
			Faculty:      "Geetha Mary",
			Day:          1,
			BuildingName: "SJT",
			StartTime:    getTime("08:00 AM"),
			EndTime:      getTime("08:50 AM"),
		},
		{
			CourseID:     "CSE1006",
			Faculty:      "Lijo VP",
			Day:          3,
			BuildingName: "SMV",
			StartTime:    getTime("10:00 AM"),
			EndTime:      getTime("10:50 AM"),
		},
		{
			CourseID:     "CSE2010",
			Faculty:      "Sasikala R",
			Day:          1,
			BuildingName: "TT",
			StartTime:    getTime("08:00 AM"),
			EndTime:      getTime("08:50 AM"),
		},
		{
			CourseID:     "CSE3045",
			Faculty:      "Raja SP",
			Day:          5,
			BuildingName: "TT",
			StartTime:    getTime("09:00 AM"),
			EndTime:      getTime("09:50 AM"),
		},
		{
			CourseID:     "CSE3054",
			Faculty:      "Abdul G",
			Day:          4,
			BuildingName: "MGR",
			StartTime:    getTime("12:00 PM"),
			EndTime:      getTime("12:50 PM"),
		},
		{
			CourseID:     "CSE2010",
			Faculty:      "Arunkumar",
			Day:          3,
			BuildingName: "SJT",
			StartTime:    getTime("10:00 AM"),
			EndTime:      getTime("10:50 AM"),
		},
	}
	db.Create(classes)
}

func getTime(timeStr string) *time.Time {
	parsedTime, err := time.Parse("03:04 PM", timeStr)
	if err != nil {
		return nil
	}
	return &parsedTime
}
