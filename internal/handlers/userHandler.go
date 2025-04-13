package handlers

import (
	"context"
	"fmt"
	"pet_project_etap1/internal/userService"
	"pet_project_etap1/internal/web/users"
)

type HandlerUser struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *HandlerUser {
	return &HandlerUser{Service: service}
}

func (u *HandlerUser) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	getAllUsers, err := u.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	responce := users.GetUsers200JSONResponse{}

	for _, usr := range getAllUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		responce = append(responce, user)
	}
	return responce, nil
}

func (u *HandlerUser) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, fmt.Errorf("error: request.Body = nil ")
	}
	userBody := request.Body

	Email := ""
	if userBody.Email != nil {
		Email = *userBody.Email
	}
	Password := ""
	if userBody.Password != nil {
		Password = *userBody.Password
	}
	userToCreate := userService.User{
		Email:    Email,
		Password: Password,
	}
	createdUser, err := u.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	responce := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password}
	return responce, nil
}

func (u *HandlerUser) UpdateUsers(_ context.Context, request users.UpdateUsersRequestObject) (users.UpdateUsersResponseObject, error) {
	if request.Body == nil {
		return nil, fmt.Errorf("error: request.Body = nil ")
	}
	userBody := request.Body
	var Id int
	Id = request.Id

	if userBody.Id != nil {
		Id = *userBody.Id
	}

	Email := ""
	if userBody.Email != nil {
		Email = *userBody.Email
	}
	Password := ""
	if userBody.Password != nil {
		Password = *userBody.Password
	}
	userToUpdate := userService.User{
		ID:       Id,
		Email:    Email,
		Password: Password,
	}
	updateUser, err := u.Service.UpdateUser(userToUpdate, request.Id)
	if err != nil {
		return nil, err
	}
	responce := users.UpdateUsers200JSONResponse{
		Email:    &updateUser.Email,
		Password: &updateUser.Password,
		Id:       &updateUser.ID}
	return responce, nil
}

func (u *HandlerUser) DeleteUsers(_ context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	err := u.Service.DeleteUser(request.Id)
	if err != nil {
		return nil, err
	}
	deleteToUsers := users.DeleteUsers204Response{}
	return deleteToUsers, nil
}
