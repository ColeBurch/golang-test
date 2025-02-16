package main

import "fmt"

// Pointers are used to reference the original value that was passed into a function

type circle struct {
	x      int
	y      int
	radius int
}

func (c *circle) grow() {
	c.radius *= 2
}

func main() {
	// You can create new pointers to a values memory address using ampersands

	myString := "hello"
	myStringPointer := &myString

	// You then use an asterisk to get to the underlying value

	fmt.Println(myStringPointer)
	fmt.Println(*myStringPointer)

	c := circle{
		x:      1,
		y:      2,
		radius: 4,
	}

	c.grow()
	fmt.Println(c.radius)
}
