package handlers

import (
	"Assignment_Golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// used mutex for manage concurrency.

func CreateEmployee(c *gin.Context) {
	var e models.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Mutex.Lock()
	e.ID = len(models.Employees) + 1
	models.Employees[e.ID] = e
	models.Mutex.Unlock()

	c.JSON(http.StatusCreated, e)
}

func GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	models.Mutex.Lock()
	employee, exists := models.Employees[id]
	models.Mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var e models.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Mutex.Lock()
	_, exists := models.Employees[id]
	if exists {
		e.ID = id
		models.Employees[id] = e
	}
	models.Mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, e)
}

func DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	models.Mutex.Lock()
	_, exists := models.Employees[id]
	if exists {
		delete(models.Employees, id)
	}
	models.Mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func GetAllEmployees(c *gin.Context) {
	models.Mutex.Lock()
	employees := make([]models.Employee, 0, len(models.Employees))
	for _, employee := range models.Employees {
		employees = append(employees, employee)
	}
	models.Mutex.Unlock()

	c.JSON(http.StatusOK, employees)
}
