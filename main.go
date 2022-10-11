package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"task-manager/controllers"
	"task-manager/database"
	"task-manager/middleware"
	"task-manager/utils"
)

func main() {
	utils.InitViper()
	host := viper.Get("HOST")
	port := viper.Get("PORT")
	address := fmt.Sprintf("%s:%s", host, port)
	database.DB = database.Connect()
	defer closeConnection()
	router := gin.Default()
	router.POST("/tasks", middleware.RequireAuth, controllers.CreateTask)
	router.POST("/users/signup", controllers.SignUpUser)
	router.POST("/users/login", controllers.LoginUser)
	router.GET("/users/validate", middleware.RequireAuth, controllers.ValidateUser)
	router.GET("/tasks/:id", middleware.RequireAuth, controllers.GetTask)
	router.GET("/tasks/createdDate/:created_at", middleware.RequireAuth, controllers.GetCreatedTask)
	router.GET("/tasks/updatedDate/:updated_at", middleware.RequireAuth, controllers.GetUpdatedTasks)
	router.GET("/tasks", middleware.RequireAuth, controllers.GetAllTasks)
	router.DELETE("/tasks/:id", middleware.RequireAuth, controllers.DeleteTask)
	router.PATCH("/tasks/:id", middleware.RequireAuth, controllers.UpdateStatus)
	router.PUT("/tasks/:id", middleware.RequireAuth, controllers.UpdateTask)

	err := router.Run(address)
	if err != nil {
		log.Fatalln("Incorrect address")
	}
}

func closeConnection() {
	sqlDB, err := database.DB.DB()
	err = sqlDB.Close()
	if err != nil {
		fmt.Println(err)
	}
}
