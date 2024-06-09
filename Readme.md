# Employee Management System

This is a simple Employee Management System built using Go and the Gin framework. It supports CRUD operations on employee records and provides a RESTful API with pagination for listing employees.

## Features

- Create a new employee
- Get details of a specific employee by ID
- Update an existing employee's details
- Delete an employee by ID
- List all employees with pagination

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Git

### Installation

1. Clone the repository

```bash
git clone " https://github.com/aamir290895/Assignment_Golang.git "
cd Assignment_Golang

2. Install dependencies

go mod tidy

go get github.com/stretchr/testify/assert

go get github.com/gin-gonic/gin


3. Run the application

go run main.go

API Endpoints


1. Create a new employee
POST /employee

Request body: 
{
    "name": "John Doe",
    "position": "Software Engineer",
    "salary": 60000
}


Response:

{
    "id": 1,
    "name": "John Doe",
    "position": "Software Engineer",
    "salary": 60000
}


Get employee details by ID
GET /employee/:id

Response:

{
    "id": 1,
    "name": "John Doe",
    "position": "Software Engineer",
    "salary": 60000
}


Update an employee's details
PUT /employee/:id

Request body:

{
    "name": "John Doe",
    "position": "Senior Software Engineer",
    "salary": 80000
}


Response:

{
    "id": 1,
    "name": "John Doe",
    "position": "Senior Software Engineer",
    "salary": 80000
}


Delete an employee by ID
DELETE /employee/:id

Response:

Status: 204 No Content
List all employees
GET /employees



List all employees
GET /employees

Response : 

[
    {
        "id": 1,
        "name": "John Doe",
        "position": "Software Engineer",
        "salary": 60000
    },
    {
        "id": 2,
        "name": "Jane Smith",
        "position": "Project Manager",
        "salary": 75000
    }
]


Concurrency Safety

The application ensures concurrency safety by using a mutex lock (sync.Mutex) around read and write operations on the in-memory employee store.


Test cases :


Explanation of Test Cases
* TestCreateEmployee: Tests the creation of an employee. It sends a POST request to create an employee and checks if the response status is 201 Created and if the employee details match.

* TestGetEmployeeByID: Tests retrieving an employee by ID. It adds an employee to the in-memory store, sends a GET request to retrieve the employee, and checks if the response matches the stored employee.

* TestUpdateEmployee: Tests updating an employee's details. It adds an employee to the store, sends a PUT request to update the employee, and checks if the updated details match the response.

* TestDeleteEmployee: Tests deleting an employee by ID. It adds an employee to the store, sends a DELETE request, and checks if the employee is removed from the store.

* TestGetAllEmployees: Tests retrieving all employees. It adds multiple employees to the store, sends a GET request to retrieve all employees, and checks if the response contains all the stored employees.


By following these steps and commands, you can thoroughly test the functionalities of your Go application and ensure that all CRUD operations work correctly.


Commands :

1. go test  // for running all the test cases.

2. go test -run [name_of_unit_test_case]  // run unit test case only

3. go test -run [name_of_unit_test_case] -v // to get vebrose output for unit test case.






