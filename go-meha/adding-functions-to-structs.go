package main

import "fmt"

type EmpIncrementStruct struct {
	name string
	salary int16
	bonus int16
}

func(emp EmpIncrementStruct) incrementedSalary(salary, bonus int16) int16 {
	return (salary + bonus)
}

func addingFuncToStruct() {
	empOne := new(EmpIncrementStruct)
	var incrementedSalary = empOne.incrementedSalary(10000, 1000)
	fmt.Println("Emp One incremented salary is: ", incrementedSalary)
}