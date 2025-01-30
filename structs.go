package main

import "fmt"

// can create structs to tell how data is structured and can be nested

type Car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	RearWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

// can also Embed structs and they inheret all the classes of the other

type Truck struct {
	Car
	TowingCapacity int
}

// Methods can be created on your struct

func (c Car) area() int {
	return c.Width * c.Height
}

// can also create anonymous structs in structs that are only expected to be used once

func main() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "Tesla",
		Model: "model 3",
	}

	fmt.Println(myCar.Make, myCar.Model)
}
