package repository

import (
    "database/sql"
    "github.com/koushik/task-reader-service/pkg/domain"
)

type TaskRepository interface {
    FetchTasks(page, limit int, status string) ([]domain.Task, error)
    FetchTaskByID(id int64) (*domain.Task, error)
}

type taskRepository struct {
    db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
    return &taskRepository{db: db}
}

func (r *taskRepository) FetchTasks(page, limit int, status string) ([]domain.Task, error) {
    offset := (page - 1) * limit
    query := `SELECT id, title, description, status, created_at, updated_at FROM tasks`
    args := []interface{}{}

    if status != "" {
        query += ` WHERE status=$1`
        args = append(args, status)
    }

    query += ` ORDER BY created_at DESC LIMIT $2 OFFSET $3`
    args = append(args, limit, offset)

    rows, err := r.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    tasks := []domain.Task{}
    for rows.Next() {
        var t domain.Task
        if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
            return nil, err
        }
        tasks = append(tasks, t)
    }

    return tasks, nil
}

func (r *taskRepository) FetchTaskByID(id int64) (*domain.Task, error) {
    query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id=$1`
    var t domain.Task
    err := r.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
    if err != nil {
        return nil, err
    }

    return &t, nil
}

