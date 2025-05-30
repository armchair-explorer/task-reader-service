package main

import (
    "database/sql"
    "log"
    "github.com/koushik/task-reader-service/pkg/api"
    "github.com/koushik/task-reader-service/pkg/repository"
    "github.com/koushik/task-reader-service/pkg/service"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/tasks_db?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewTaskRepository(db)
    svc := service.NewTaskService(repo)
    handler := api.NewTaskHandler(svc)

    r := gin.Default()

    r.GET("/tasks", handler.GetTasks)
    r.GET("/tasks/:id", handler.GetTaskByID)

    r.Run(":8070")
}

