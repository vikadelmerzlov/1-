package taskService

import (
	"fmt"
	"gorm.io/gorm"
	"pet_project_etap1/internal/web/users"
)

type TaskRepository interface {
	GetAllTasks() ([]Task, error)
	//Возвращаем массив из всех бд, и ошибку
	CreateTask(task Task) (Task, error)
	//Передаем в функцию task типа Task, и возвращаем созданный Task и err
	UpdateTask(id int, task Task) (Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task и err
	DeleteTask(id int) error
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	GetTasksByUserID(userID int) ([]Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByUserID(userID int) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_ID=?", userID).Find(&tasks).Error
	fmt.Println("error finding tasks by user id:", userID)
	return tasks, err
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	var user users.User
	if err := r.db.First(&user, task.UserID).Error; err != nil {
		fmt.Println("Error creating task", task)
		return Task{}, err
	}
	if err := r.db.Create(&task).Error; err != nil {
		fmt.Println("error creating tasks", task)
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) UpdateTask(id int, task Task) (Task, error) {
	if err := r.db.Model(&Task{}).Where("id=?", id).Update("description", task.Description).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) DeleteTask(id int) error {
	if err := r.db.Delete(&Task{}, id); err != nil {
		return err.Error
	}
	return nil
}
