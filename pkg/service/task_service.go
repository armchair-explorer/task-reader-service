package service

import (
    "github.com/koushik/task-reader-service/pkg/domain"
    "github.com/koushik/task-reader-service/pkg/repository"
)

type TaskService interface {
    GetTasks(page, limit int, status string) ([]domain.Task, error)
    GetTaskByID(id int64) (*domain.Task, error)
}

type taskService struct {
    repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
    return &taskService{repo: repo}
}

func (s *taskService) GetTasks(page, limit int, status string) ([]domain.Task, error) {
    return s.repo.FetchTasks(page, limit, status)
}

func (s *taskService) GetTaskByID(id int64) (*domain.Task, error) {
    return s.repo.FetchTaskByID(id)
}

