package main

import "fmt"

// In go functions can be passed as data

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(2, 3, 4, add))
	fmt.Println(aggregate(2, 3, 4, multiply))
}
