package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	slice := numbers[:3]

	slice = append(slice, 100)

	fmt.Println(numbers)
	fmt.Println(slice)
	fmt.Println(len(numbers))
	fmt.Println(len(slice))

	newNumbers := []int{10, 20, 30}
	newSlice := newNumbers

	newSlice = append(newSlice, 100)

	fmt.Println(newNumbers)
	fmt.Println(newSlice)
	fmt.Println(len(newNumbers))
	fmt.Println(len(newSlice))
}
