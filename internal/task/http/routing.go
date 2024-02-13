package http

import (
	"github.com/gin-gonic/gin"
	"github.com/elvinlari/docker-golang/cmd/middleware"
)

// TaskServiceHTTPHandlers defines all the handlers the TaskService needs. It's
// possible to register routes for a different implementation (like a mock).
type TaskServiceHTTPHandlers interface {
	GetTask(c *gin.Context)
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	DeleteTasks(c *gin.Context)
}

func RegisterRoutes(r *gin.Engine, t TaskServiceHTTPHandlers) {
    tasks := r.Group("/tasks")
    {
        tasks.GET("/", middleware.CommonHeaders(t.GetTasks))
        tasks.GET("/:taskid", middleware.CommonHeaders(t.GetTask))
        tasks.POST("/", middleware.IsAuthorizedJWT(t.CreateTask, "pet-details"))
        tasks.DELETE("/:taskid", middleware.IsAuthorizedJWT(t.DeleteTask, "pet-details"))
        tasks.DELETE("/", middleware.IsAuthorizedJWT(t.DeleteTasks, "pet-details"))
    }
}

