package main

import (
	"fmt"
)

// 实现一个计算器的程序
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

type OperationFactory struct{}

func (factory *OperationFactory) GetAddOperation() Operation {
	return new(AddOperation)
}

func (factory *OperationFactory) GetSubOperation() Operation {
	return new(SubOperation)
}
func (factory *OperationFactory) GetMulOperation() Operation {
	return new(MulOperation)
}
func (factory *OperationFactory) GetDivOperation() Operation {
	return new(DivOperation)
}
func main() {
	factory := new(OperationFactory)
	num1, num2 := 100, 90

	operation := factory.GetAddOperation()
	fmt.Printf("%v + %v = %v \n", num1, num2, operation.Do(num1, num2))

	operation = factory.GetSubOperation()
	fmt.Printf("%v - %v = %v \n", num1, num2, operation.Do(num1, num2))

	operation = factory.GetMulOperation()
	fmt.Printf("%v * %v = %v \n", num1, num2, operation.Do(num1, num2))

	operation = factory.GetDivOperation()
	fmt.Printf("%v / %v = %v \n", num1, num2, operation.Do(num1, num2))

	fmt.Println("hi,design pattern go implements.")
}
