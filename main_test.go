package main

import (
	"assignment_golang/handlers"
	"assignment_golang/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/employee", handlers.CreateEmployee)
	router.GET("/employee/:id", handlers.GetEmployeeByID)
	router.PUT("/employee/:id", handlers.UpdateEmployee)
	router.DELETE("/employee/:id", handlers.DeleteEmployee)
	router.GET("/employees", handlers.GetAllEmployees)
	return router
}

func TestCreateEmployee(t *testing.T) {
	t.Log("Starting Test")

	router := setupRouter()

	employee := models.Employee{
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   60000,
	}
	jsonValue, _ := json.Marshal(employee)

	req, _ := http.NewRequest("POST", "/employee", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdEmployee models.Employee
	json.Unmarshal(w.Body.Bytes(), &createdEmployee)

	assert.Equal(t, employee.Name, createdEmployee.Name)
	assert.Equal(t, employee.Position, createdEmployee.Position)
	assert.Equal(t, employee.Salary, createdEmployee.Salary)

	t.Log(createdEmployee)
	t.Log("Finished Test")

}

func TestGetEmployeeByID(t *testing.T) {

	t.Log("Starting Test")

	router := setupRouter()

	// Add an employee to the store for testing
	models.Mutex.Lock()
	employee := models.Employee{
		ID:       1,
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   60000,
	}
	models.Employees[1] = employee
	models.Mutex.Unlock()

	req, _ := http.NewRequest("GET", "/employee/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var retrievedEmployee models.Employee
	json.Unmarshal(w.Body.Bytes(), &retrievedEmployee)

	assert.Equal(t, employee.ID, retrievedEmployee.ID)
	assert.Equal(t, employee.Name, retrievedEmployee.Name)
	assert.Equal(t, employee.Position, retrievedEmployee.Position)
	assert.Equal(t, employee.Salary, retrievedEmployee.Salary)

	t.Log(retrievedEmployee)
	t.Log("Finished Test")

}

func TestUpdateEmployee(t *testing.T) {

	t.Log("Starting Test")

	router := setupRouter()

	// Add an employee to the store for

	models.Mutex.Lock()
	employee := models.Employee{
		ID:       1,
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   60000,
	}
	models.Employees[1] = employee
	models.Mutex.Unlock()

	updatedEmployee := models.Employee{
		Name:     "John Doe Updated",
		Position: "Senior Software Engineer",
		Salary:   80000,
	}
	jsonValue, _ := json.Marshal(updatedEmployee)

	req, _ := http.NewRequest("PUT", "/employee/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseEmployee models.Employee
	json.Unmarshal(w.Body.Bytes(), &responseEmployee)

	assert.Equal(t, updatedEmployee.Name, responseEmployee.Name)
	assert.Equal(t, updatedEmployee.Position, responseEmployee.Position)
	assert.Equal(t, updatedEmployee.Salary, responseEmployee.Salary)

	t.Log(responseEmployee)
	t.Log("Finished Test")

}

func TestDeleteEmployee(t *testing.T) {

	t.Log("Starting Test")

	router := setupRouter()

	// Add an employee to the store for testing
	models.Mutex.Lock()
	employee := models.Employee{
		ID:       1,
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   60000,
	}
	models.Employees[1] = employee
	models.Mutex.Unlock()

	req, _ := http.NewRequest("DELETE", "/employee/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	models.Mutex.Lock()
	_, exists := models.Employees[1]
	models.Mutex.Unlock()

	assert.False(t, exists)
	t.Log("Finished Test")

}

func TestGetAllEmployees(t *testing.T) {
	t.Log("Starting Test")

	router := setupRouter()

	// Add some employees to the store for testing
	models.Mutex.Lock()
	models.Employees[1] = models.Employee{ID: 1, Name: "John Doe", Position: "Software Engineer", Salary: 60000}
	models.Employees[2] = models.Employee{ID: 2, Name: "Jane Smith", Position: "Project Manager", Salary: 75000}
	models.Mutex.Unlock()

	req, _ := http.NewRequest("GET", "/employees", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var employees []models.Employee
	json.Unmarshal(w.Body.Bytes(), &employees)

	assert.Len(t, employees, 2)

	t.Log(employees)
	t.Log("Finished Test")

}
