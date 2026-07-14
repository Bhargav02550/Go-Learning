package main

import (
	"fmt"
)

type Student struct {
	Id    int
	Name  string
	Marks int
}

func DemoMaps() {

	students := map[int]Student{
		1: {
			Id:    1,
			Name:  "Bhargav",
			Marks: 100,
		},
		2: {
			Id:    2,
			Name:  "Pavan",
			Marks: 99,
		},
		3: {
			Id:    3,
			Name:  "Kumar",
			Marks: 98,
		},
	}

	for key, value := range students {
		fmt.Println(key, value)
	}

	student, found := students[1]

	if found {
		fmt.Printf("Student found with id : %d", student.Id)
	} else {
		fmt.Printf("Student Not found with id : %d", student.Id)
	}

	updateStudent := students[1]

	updateStudent.Marks = 99

	students[1] = updateStudent

	delete(students, 3)

	for key, value := range students {
		fmt.Println(key, value)
	}
}

func main() {
	DemoMaps()
}
