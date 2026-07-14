package main

import "fmt"

func main() {
	employee := NewEmployee(1, "Bhargav", 35000.00)

	employee.PrintDetails()

	employee.IncreaseSalary(10)

	fmt.Println()

	employee.PrintDetails()

	fmt.Printf("Annual Salary: %2f\n", employee.AnnualSalary())

}
