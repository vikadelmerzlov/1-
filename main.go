package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	Is_Done     bool   `json:"is_Done"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных :%v", err)
	}
	db.AutoMigrate(&Task{})
}

func GetHandler(c echo.Context) error {
	var taskes []Task
	if err := db.Find(&taskes).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "The message was not delivered ",
		})
	}
	return c.JSON(http.StatusOK, &taskes)
}

func PostHandler(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "The message was not received",
		})
	}

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error ",
			Message: "The message was not updated",
		})

	}
	return c.JSON(http.StatusOK, Response{
		Status:  "Ok",
		Message: "The message was successful update",
	})
}

func main() {
	initDB()
	e := echo.New()
	e.GET("/api/task", GetHandler)
	e.POST("/api/task", PostHandler)
	e.Start("localhost:8080")
}
