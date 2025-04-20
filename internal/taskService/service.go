package taskService

import "fmt"

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	fmt.Println("task in service", task)
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTasksByUserID(userID int) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)

}

func (s *TaskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTasks()

}

func (s *TaskService) UpdateTask(id int, task Task) (Task, error) {
	return s.repo.UpdateTask(id, task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)

}
