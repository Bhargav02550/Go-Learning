package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float64
}

func NewEmployee(id int, name string, salary float64) *Employee {
	return &Employee{
		Id:     id,
		Name:   name,
		Salary: salary,
	}
}

func (e *Employee) PrintDetails() {
	fmt.Printf("ID: %d\n", e.Id)
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Salary: %.2f\n", e.Salary)
}

func (e *Employee) IncreaseSalary(percent float64) {
	e.Salary += e.Salary * (percent / 100)
}

func (e *Employee) AnnualSalary() float64 {
	return e.Salary * 12
}
