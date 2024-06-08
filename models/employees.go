package models

import "sync"

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

var (
	Employees = make(map[int]Employee)
	Mutex     = &sync.Mutex{}
)
