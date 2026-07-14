package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 5, 10)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers2 := make([]int, 0, 5)

	numbers2 = append(numbers2, 100)
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))

	numbers2 = append(numbers2, 200)
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))

	numbers2 = append(numbers2, 300)
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))

	numbers2 = append(numbers2, 400)
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))

	numbers2 = append(numbers2, 500)
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))

	numbers2 = append(numbers2, 600)
	fmt.Println(numbers2)
	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))

	original := []int{1, 2, 3, 4, 5}

	newCopy := make([]int, len(original))

	copy(newCopy, original)

	fmt.Println(newCopy)
}
