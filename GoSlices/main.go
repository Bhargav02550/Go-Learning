package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30, 40, 50}

	slice1 := numbers[:3]
	// Print the slice
	fmt.Print(slice1)

	// Print len(numbers)
	fmt.Print(len(numbers))

	// Print cap(numbers)
	fmt.Print(cap(numbers))

	// Create a slice from index 1 to 4
	slice2 := numbers[1:4]

	// Print the new slice
	fmt.Print(slice2)

	// Print its length
	fmt.Print(len(slice2))

	// Print its capacity
	fmt.Print(cap(slice2))

	// Change the first element of the new slice to 999
	slice2[0] = 999

	// Print the original slice
	for _, value := range numbers {
		fmt.Println(value)
	}

	// Print the new slice
	for _, value := range slice2 {
		fmt.Println(value)
	}

}
