package userService

import "pet_project_etap1/internal/taskService"

type UserService struct {
	repo UserRepository
}

func NewService(repo *userRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUsers() ([]User, error) {
	return u.repo.GetUsers()
}

func (u *UserService) GetTasksByUserID(userId int) ([]taskService.Task, error) {
	return u.repo.GetTasksByUserID(userId)
}

func (u *UserService) CreateUser(user User) (User, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) UpdateUser(user User, id int) (User, error) {
	return u.repo.UpdateUser(user, id)
}

func (u *UserService) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
