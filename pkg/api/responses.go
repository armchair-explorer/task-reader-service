package api

import "github.com/koushik/task-reader-service/pkg/domain"

type PaginatedTasks struct {
    Tasks []domain.Task `json:"tasks"`
    Page  int           `json:"page"`
    Limit int           `json:"limit"`
}

