package main

import (
	"errors"
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

func (factory *OperationFactory) GetOperation(opt string) (Operation, error) {
	switch opt {
	case "+":
		return new(AddOperation), nil
	case "-":
		return new(SubOperation), nil
	case "*":
		return new(MulOperation), nil
	case "/":
		return new(DivOperation), nil
	default:
		return nil, errors.New("unsupport opt type symbol")
	}
}

func main() {
	factory := new(OperationFactory)

	num1, num2 := 1, 2
	opt := "+"
	operation, err := factory.GetOperation(opt)
	if err != nil {
		fmt.Errorf("err:%v,opt:%v", err, opt)
	}
	fmt.Printf("%v %v %v = %v \n", num1, opt, num2, operation.Do(num1, num2))

	opt = "-"
	operation, err = factory.GetOperation(opt)
	if err != nil {
		fmt.Errorf("err:%v,opt:%v", err, opt)
	}
	fmt.Printf("%v %v %v = %v \n", num1, opt, num2, operation.Do(num1, num2))

	opt = "*"
	operation, err = factory.GetOperation(opt)
	if err != nil {
		fmt.Errorf("err:%v,opt:%v", err, opt)
	}
	fmt.Printf("%v %v %v = %v \n", num1, opt, num2, operation.Do(num1, num2))

	opt = "/"
	operation, err = factory.GetOperation(opt)
	if err != nil {
		fmt.Errorf("err:%v,opt:%v", err, opt)
	}
	fmt.Printf("%v %v %v = %v \n", num1, opt, num2, operation.Do(num1, num2))

	fmt.Println("hi,design pattern go implements.")
}
