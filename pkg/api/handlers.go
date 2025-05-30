package api

import (
    "net/http"
    "strconv"
    "github.com/koushik/task-reader-service/pkg/service"

    "github.com/gin-gonic/gin"
)

type TaskHandler struct {
    service service.TaskService
}

func NewTaskHandler(s service.TaskService) *TaskHandler {
    return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    status := c.Query("status")

    tasks, err := h.service.GetTasks(page, limit, status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
        return
    }

    task, err := h.service.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

