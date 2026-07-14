package main

import (
	"fmt"
)

func main() {

	//Creating an array in Go this means Create an array of size 5 with type int
	numbers := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	//Accessing Numbers from Array
	fmt.Println(numbers[0])
	fmt.Println(numbers[3])

	//Length of array
	fmt.Println(len(numbers))

	numbers2 := numbers

	numbers2[0] = 100

	for i := 0; i < len(numbers2); i++ {
		fmt.Println(numbers2[i])
	}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// _ means it does not bother about the value in that place
	for _, value := range numbers {
		fmt.Println(value)
	}
}
