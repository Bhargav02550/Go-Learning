package main

import (
	"fmt"
)

func DemoMapExamples() {
	//Dictionary in Go
	//map[KeyType]ValueType

	//Using Make
	Dict := make(map[int]string)

	Dict[1] = "Testing1"
	Dict[2] = "Testing2"

	//Using Literal
	students := map[int]string{
		1: "Bhargav",
		2: "Pavan",
		3: "Kumar",
	}

	fmt.Println(students)

	//Accessing Values
	fmt.Println(students[2])

	//Go Prints "" if the key does not exist
	//This is because the default value of String is ""
	//Applies for all the types, It just prints Defaults

	//Checking existance of Value using Key
	name, exists := students[1]

	fmt.Println(name)
	fmt.Println(exists)

	// Updating Values
	// Replaces the existing ones now Value for Key 1 becomes New Bhargav
	students[1] = "New Bhargav"

	fmt.Println(students)

	// Deleting Values
	// Syntax : delete(map , key)
	delete(students, 2)

	fmt.Println(students)

	// Len of Map
	fmt.Println(len(students))

	// Iterating over the map
	for key, value := range students {
		fmt.Println(key, value)
	}

	//Map of Structs

	type Employee struct {
		Id   int
		Name string
	}

	employees := make(map[int]Employee)

	employees[1] = Employee{
		Id:   1,
		Name: "Bhargav",
	}

	employee := employees[1]
	fmt.Println(employee.Id)
	fmt.Println(employees[1].Name)

	//Nested Maps
	marks := map[string]map[string]int{
		"Bhargav": {
			"Math": 99,
		},
	}

	fmt.Println(marks["Bhargav"]["Math"])

	//Interview Questions
	//var students map[int]string
	//students[1] = "Bhargav"
	// Unlike Slices Maps need to be initialized first,
	// We cannot write to nil maps
}
