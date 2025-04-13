package userService

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUsers() ([]User, error) {
	return u.repo.GetUsers()
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
