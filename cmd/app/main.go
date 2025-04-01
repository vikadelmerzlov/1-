package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet_project_etap1/internal/database"
	"pet_project_etap1/internal/handlers"
	"pet_project_etap1/internal/taskService"
	"pet_project_etap1/internal/web/tasks"
)

/*
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

	func getHandler(c echo.Context) error {
		var taskes []Task
		if err := db.Find(&taskes).Error; err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "The message was not delivered ",
			})
		}
		return c.JSON(http.StatusOK, &taskes)
	}

	func postHandler(c echo.Context) error {
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
		return c.JSON(http.StatusOK, &task)
	}

	func patchHandler(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "Incorrect ID",
			})
		}

		var updateTask Task
		if err := c.Bind(&updateTask); err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "Invalid input",
			})
		}

		if err := db.Model(&Task{}).Where("id=?", id).Update("description", updateTask.Description).Error; err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "The message was not updated",
			})
		}

		updateTask.ID = id
		return c.JSON(http.StatusOK, &updateTask)
	}

	func deleteHandler(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "Incorrect ID",
			})
		}

		if err := db.Delete(&Task{}, id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "The message was not deleted",
			})
		}

		return c.NoContent(http.StatusOK)
	}
*/
func main() {
	database.InitDB()

	e := echo.New()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8040"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
