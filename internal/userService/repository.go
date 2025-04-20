package userService

import (
	"gorm.io/gorm"
	"pet_project_etap1/internal/taskService"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User, id int) (User, error)
	DeleteUser(id int) error
	GetTasksByUserID(userID int) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUsers() ([]User, error) {
	var users []User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *userRepository) GetTasksByUserID(userID int) ([]taskService.Task, error) {
	var task []taskService.Task
	err := ur.db.Where("id = ?", userID).Find(&task).Error
	return task, err

}

func (ur *userRepository) CreateUser(user User) (User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(user User, id int) (User, error) {
	if err := ur.db.Model(&User{}).Where("id = ?", id).Update("email", user.Email).Update("password", user.Password).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(id int) error {
	if err := ur.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
