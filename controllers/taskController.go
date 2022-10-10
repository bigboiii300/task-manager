package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-manager/database"
	"task-manager/models"
)

func CreateTask(c *gin.Context) {
	newTask := models.Task{}
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	if newTask.Status == "todo" || newTask.Status == "in progress" || newTask.Status == "done" {
		database.DB.Create(&newTask)
		c.IndentedJSON(http.StatusCreated, newTask)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := database.DB.First(&task, c.Param("id")).Delete(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, tasks)
	}
}

func GetTask(c *gin.Context) {
	var task models.Task
	fmt.Println(c.Param("id"))
	if err := database.DB.First(&task, c.Param("id")).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

func UpdateStatus(c *gin.Context) {
	var task models.Task
	if err := database.DB.First(&task, c.Param("id")).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// Validate input
	var input models.UpdateStatus
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Status == "todo" || input.Status == "in progress" || input.Status == "done" {
		database.DB.Model(&task).Update("status", input.Status)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := database.DB.First(&task, c.Param("id")).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// Validate input
	var input models.UpdateTask
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&task).Update("task_name", input.TaskName)
	database.DB.Model(&task).Update("task_details", input.TaskDetails)
	database.DB.Model(&task).Update("completion_date", input.CompletionDate)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetUpdatedTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Raw("SELECT * FROM tasks WHERE date(updated_at) = ?", c.Param("updated_at")).Scan(&tasks)
	if len(tasks) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, tasks)
	}
}

func GetCreatedTask(c *gin.Context) {
	var tasks []models.Task
	database.DB.Raw("SELECT * FROM tasks WHERE date(created_at) = ?", c.Param("created_at")).Scan(&tasks)
	if len(tasks) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, tasks)
	}
}
