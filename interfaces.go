// interfaces are a collection of functions that implement a common way
// to interact with multiple structs
package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type W2 struct {
	name string
	yearlySalary int
}

func (e W2) getName() string {
	return e.name
}

func (e W2) getSalary() int {
	return e.yearlySalary
}

type hourlyEmployee struct {
	name string
	hourlySalary int
	hoursPerYear int
}

func (e hourlyEmployee) getName() string {
	return e.name
}

func (e hourlyEmployee) getSalary() int {
	return e.hourlySalary * e.hoursPerYear
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
}

func main() {
	test(W2{
		name: "Jack",
		yearlySalary: 50000,
	})
	test(hourlyEmployee{
		name: "Bob",
		hourlySalary: 100,
		hoursPerYear: 73,
	})
}
