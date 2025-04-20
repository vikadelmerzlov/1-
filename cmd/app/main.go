package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet_project_etap1/internal/database"
	"pet_project_etap1/internal/handlers"

	"pet_project_etap1/internal/taskService"
	"pet_project_etap1/internal/userService"
	"pet_project_etap1/internal/web/tasks"
	"pet_project_etap1/internal/web/users"
)

func main() {

	database.InitDB()

	e := echo.New()

	repoTask := taskService.NewTaskRepository(database.DB)
	serviceTask := taskService.NewService(repoTask)

	taskHandler := handlers.NewTaskHandler(serviceTask)

	repoUser := userService.NewUserRepository(database.DB)
	serviceUser := userService.NewService(repoUser)
	userHandler := handlers.NewUserHandler(serviceUser)

	e.Use(middleware.Logger())

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1024, // Размер стека для вывода
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			log.Printf("Panic recovered: %v\nStack trace:\n%s", err, stack)
			return nil // Возвращаем nil, чтобы echo не генерировал HTTP-ответ
		},
	}))

	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8040"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

}
