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






