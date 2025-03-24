package handlers

import (
	"context"
	"pet_project_etap1/internal/taskService"
	"pet_project_etap1/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
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
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	getAllTasks, err := h.Service.GetAllTask()
	if err != nil {
		return nil, err
	}
	responce := tasks.GetTasks200JSONResponse{}

	for _, tsk := range getAllTasks {
		task := tasks.Task{
			Description: &tsk.Description,
			Id:          &tsk.ID,
			IsDone:      &tsk.IsDone,
			Title:       &tsk.Title,
		}
		responce = append(responce, task)
	}
	return responce, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		IsDone:      *taskRequest.IsDone,
		Title:       *taskRequest.Title,
		Description: *taskRequest.Description,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:          &createdTask.ID,
		IsDone:      &createdTask.IsDone,
		Title:       &createdTask.Title,
		Description: &createdTask.Description,
	}
	return response, nil
}

/*func (h *Handler) GetTasksHandler(c echo.Context) error {
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
}

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

func (h *Handler) DeleteTask(c echo.Context) error {
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
*/
