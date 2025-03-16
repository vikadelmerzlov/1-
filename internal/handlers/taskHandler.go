package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pet_project_etap1/internal/taskService"
	"strconv"
)

var nextID = 1

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

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
	return c.NoContent(http.StatusOK)
}
