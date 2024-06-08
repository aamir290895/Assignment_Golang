package main

import (
	"Assignment_Golang/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/employees", handlers.GetAllEmployees)
	r.POST("/create", handlers.CreateEmployee)
	r.GET("/employee/:id", handlers.GetEmployeeByID)
	r.PUT("/employee/:id", handlers.UpdateEmployee)
	r.DELETE("/delete/:id", handlers.DeleteEmployee)

	r.Run(":8080")

}
