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

func (h *HandlerTask) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, fmt.Errorf("error: request.Body = nil ")
	}
	taskRequest := request.Body
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

	taskToCreate := taskService.Task{
		Is_Done:     IsDone,
		Title:       Title,
		Description: Description,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
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

	/*responce := tasks.UpdateTask200JSONResponse{
	//Id:          &updatedTask.ID,
	Description: &updatedTask.Description,
	/*IsDone:      &updatedTask.Is_Done,
	Title:       &updatedTask.Title,*/
	//}
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

/*
var nextID int

	func (h *Handler) GetTasksHandler(c echo.Context) error {
		tasks, err := h.Service.GetAllTask()
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.JSON(http.StatusOK, &tasks)
	}

	func (h *Handler) CreateTask(c echo.Context) error {
		var task taskService.Task
		if err := c.Bind(&task); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		createdTask, err := h.Service.CreateTask(task)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		createdTask = task
		createdTask.ID = nextID
		nextID++

		return c.JSON(http.StatusOK, createdTask)

	func (h *Handler) UpdateTask(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		var updateTask taskService.Task
		if err := c.Bind(&updateTask); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		updateTsk, err := h.Service.UpdateTask(id, updateTask)

		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		updateTsk.ID = id
		return c.JSON(http.StatusOK, &updateTsk)

}
*/
/*func (h *Handler) DeleteTask(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = h.Service.DeleteTask(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusNoContent)
}

/*
	func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
		//TODO implement me+
		panic("implement me")
	}

	type Task struct {
		ID          int    `json:"id"`
		Is_Done     bool   `json:"is_Done"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	func (h *Handler) PostTasks(_ context.Context, _ tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
		//TODO implement me
		panic("implement me")
	}
*/
