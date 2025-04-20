package handlers

import (
	"context"
	"fmt"
	"log"
	"pet_project_etap1/internal/taskService"
	"pet_project_etap1/internal/web/tasks"
)

type HandlerTask struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *HandlerTask {
	return &HandlerTask{Service: service}
}

func (h *HandlerTask) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	getAllTasks, err := h.Service.GetAllTask()
	if err != nil {
		return nil, err
	}
	responce := tasks.GetTasks200JSONResponse{}

	for _, tsk := range getAllTasks {
		task := tasks.Task{
			Description: &tsk.Description,
			Id:          &tsk.ID,
			IsDone:      &tsk.Is_Done,
			Title:       &tsk.Title,
		}
		responce = append(responce, task)
	}
	return responce, nil
}

func (h *HandlerTask) GetTasksByUserID(_ context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	userID := request.UserId // Получаем ID из параметров URL
	tasksByUserID, err := h.Service.GetTasksByUserID(userID)
	if err != nil {
		return nil, err // Важно обрабатывать ошибки
	}

	responce := tasks.GetTasksByUserID200JSONResponse{}
	for _, tsk := range tasksByUserID {
		task := tasks.Task{
			Description: &tsk.Description,
			Id:          &tsk.ID,
			IsDone:      &tsk.Is_Done,
			Title:       &tsk.Title,
			UserId:      &tsk.UserID,
		}
		responce = append(responce, task)
	}

	return responce, nil
}

func (h *HandlerTask) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, fmt.Errorf("error: request.Body = nil")
	}
	taskRequest := request.Body
	//userID := *taskRequest.UserId
	userID := 0
	if taskRequest.UserId != nil {
		userID = *taskRequest.UserId
		fmt.Printf("userID: %d\n", userID)
	}
	IsDone := false
	if taskRequest.IsDone != nil {
		IsDone = *taskRequest.IsDone
		fmt.Printf("IsDone: %v\n", IsDone)
	}

	Title := ""
	if taskRequest.Title != nil {
		Title = *taskRequest.Title
		fmt.Printf("Title: %s\n", Title)
	}
	Description := ""
	if taskRequest.Description != nil {
		Description = *taskRequest.Description
		fmt.Printf("Description: %s\n", Description)
	}

	taskToCreate := taskService.Task{
		Is_Done:     IsDone,
		Title:       Title,
		Description: Description,
		UserID:      userID,
	}
	fmt.Println(taskToCreate)
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		// **Логирование ошибки!**
		log.Printf("Error creating task in service: %v\n", err)
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		UserId:      &createdTask.UserID,
		Id:          &createdTask.ID,
		IsDone:      &createdTask.Is_Done,
		Title:       &createdTask.Title,
		Description: &createdTask.Description,
	}
	return response, nil
}

func (h *HandlerTask) UpdateTasks(_ context.Context, request tasks.UpdateTasksRequestObject) (tasks.UpdateTasksResponseObject, error) {
	log.Println("Запрос на обновление задачи получен")
	if request.Body == nil {
		log.Println("error: request.Body = nil ")
		return nil, fmt.Errorf("error: request.Body = nil ")
	}
	taskRequest := request.Body

	var Id int

	if taskRequest.Id != nil {
		Id = *taskRequest.Id
		log.Printf("Id: задачи для обновления %d", Id)
	} else {
		log.Println("error : id отсутствует")
	}
	Id = request.Id

	IsDone := false
	if taskRequest.IsDone != nil {
		IsDone = *taskRequest.IsDone
	}
	Title := ""
	if taskRequest.Title != nil {
		Title = *taskRequest.Title
	}
	Description := ""
	if taskRequest.Description != nil {
		Description = *taskRequest.Description
	}

	taskToUpdate := taskService.Task{
		Is_Done:     IsDone,
		Title:       Title,
		Description: Description,
		ID:          Id,
	}

	updatedTask, err := h.Service.UpdateTask(request.Id, taskToUpdate)
	if err != nil {
		return nil, err
	}

	responce := tasks.UpdateTasks200JSONResponse{
		Description: &updatedTask.Description,
		Id:          &updatedTask.ID,
		IsDone:      &updatedTask.Is_Done,
		Title:       &updatedTask.Title,
	}

	return responce, nil

}

func (h *HandlerTask) DeleteTasks(_ context.Context, request tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	err := h.Service.DeleteTask(request.Id)
	if err != nil {
		return nil, err
	}
	deleteToTask := tasks.DeleteTasks204Response{}
	return deleteToTask, nil
}
