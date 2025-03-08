package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Task struct {
	ID      int  `json:"id"`
	IS_Done bool `json:"is_Done"`
}

type Responce struct {
	Task string `json:"task"`
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
		return c.JSON(http.StatusBadRequest, Responce{Task: "Error"})
	}
	return c.JSON(http.StatusOK, &taskes)
}

func PostHandler(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, Responce{Task: "error"})
	}

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Responce{Task: "error"})

	}
	return c.JSON(http.StatusOK, Responce{Task: "Ok"})
}

func main() {
	initDB()
	e := echo.New()
	e.GET("/api", GetHandler)
	e.POST("/api", PostHandler)
	e.Start("localhost:8080")
}
