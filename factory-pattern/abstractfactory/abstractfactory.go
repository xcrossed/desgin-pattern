package main

import (
	"fmt"
)

type Factory interface {
	Create() Operation
}

type Operation interface {
	Do(num1 int, num2 int) int
}

type AddOperation struct{}

func (add *AddOperation) Do(num1, num2 int) int {
	return num1 + num2
}

type SubOperation struct{}

func (add *SubOperation) Do(num1, num2 int) int {
	return num1 - num2
}

type MulOperation struct{}

func (add *MulOperation) Do(num1, num2 int) int {
	return num1 * num2
}

type DivOperation struct{}

func (add *DivOperation) Do(num1, num2 int) int {
	return num1 / num2
}

type AddFactory struct{}

func (factory *AddFactory) Create() Operation {
	return new(AddOperation)
}

type SubFactory struct{}

func (factory *SubFactory) Create() Operation {
	return new(SubOperation)
}

type MulFactory struct{}

func (factory *MulFactory) Create() Operation {
	return new(MulOperation)
}

type DivFactory struct{}

func (factory *DivFactory) Create() Operation {
	return new(DivOperation)
}

func main() {
	var factory Factory
	var operation Operation
	num1, num2 := 100, 90

	factory = new(AddFactory)
	operation = factory.Create()
	fmt.Printf("%v + %v = %v \n", num1, num2, operation.Do(num1, num2))

	factory = new(SubFactory)
	operation = factory.Create()
	fmt.Printf("%v - %v = %v \n", num1, num2, operation.Do(num1, num2))

	factory = new(MulFactory)
	operation = factory.Create()
	fmt.Printf("%v * %v = %v \n", num1, num2, operation.Do(num1, num2))

	factory = new(DivFactory)
	operation = factory.Create()
	fmt.Printf("%v / %v = %v \n", num1, num2, operation.Do(num1, num2))

	fmt.Println("hi,design pattern go implements.")
}
