package main

func main() {

	// arrays in go are fixed size and you put there size in the intialization

	primeArray := [6]int{2, 3, 5, 7, 11, 13}

	// Slices are portions of arrays that are defined with left index including
	// and right index not included. Slices are a reference to an underlying
	// array so when you modify a slice it modifies the original array even
	// if it was not defined within the same function

	primeSlice := primeArray[1:4]

	// Because arrays are fixed in size you often work more with slices as
	// they can grow and shrink as needed. Arrays are stored in sequence in
	// memory so if a slice needs to grow past the arrays size it will allocate
	// a new postion in memory for the array to go

	// Slices can be created using the make function where the syntax is
	// make([]type, initial length of the array, capacity in memory)

	mySlice := make([]int, 5, 10)

	// Due to ease of use, unless you are having performance issues most slices
	// are written without the capacity number and the capacity defaults to the
	// length of the slice

	newSlice := make([]int, 5)

	// You can also use slice literals

	literalSlice := []string{"I", "am", "defined"}

	// len() and cap() are built in functions to tell you how many items are in
	// a slice and the capacity of that slice

	// Use the append function to append items to the end of a slice

	newSlice = append(newSlice, 1)

}
