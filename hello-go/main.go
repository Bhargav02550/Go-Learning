package main

import "fmt"

func main() {
	a := 10
	b := 5

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Sum        :%d\n", Add(a, b))
	fmt.Printf("Difference :%d\n", Diff(a, b))
	fmt.Printf("Product    :%d\n", Multiply(a, b))
	quot, err := Quot(a, b)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Quotient   :%d\n", quot)
	}
	fmt.Printf("Reminder   :%d\n", Remainder(a, b))
}

func Add(a, b int) int {
	return a + b
}

func Diff(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Quot(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division By Zero")
	}
	return b / a, nil
}

func Remainder(a, b int) int {
	return b % a
}
