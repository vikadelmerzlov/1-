package userService

import "pet_project_etap1/internal/taskService"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []taskService.Task
}
