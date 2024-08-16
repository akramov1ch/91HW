package main

import (
	"91HW/config"
	"91HW/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := config.SetupMongoDB()
	handler := handlers.NewTaskHandler(db)

	router.POST("/tasks", handler.CreateTask)
	router.GET("/tasks", handler.GetTasks)
	router.GET("/tasks/:id", handler.GetTaskByID)
	router.PUT("/tasks/:id", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.DeleteTask)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Serverni ishga tushirishda xatolik yuz berdi: %v", err)
	}
}
