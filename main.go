package main

import (
	"assignment_golang/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/employees", handlers.GetAllEmployees)
	r.POST("/create", handlers.CreateEmployee)
	r.GET("/employee/:id", handlers.GetEmployeeByID)
	r.PUT("/employee/:id", handlers.UpdateEmployee)
	r.DELETE("/delete/:id", handlers.DeleteEmployee)

	r.Run(":8080")

}
