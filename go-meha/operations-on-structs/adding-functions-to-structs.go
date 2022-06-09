package main

import "fmt"

type EmpIncrementStruct struct {
	name   string
	salary int16
	bonus  int16
}

func (emp *EmpIncrementStruct) incrementedSalary() int16 {
	return (emp.salary + emp.bonus)
}

func addingFuncToStruct() {
	empOne := new(EmpIncrementStruct)
	empOne.salary = 2000
	empOne.bonus = 100
	var incrementedSalary = empOne.incrementedSalary()
	fmt.Println("Emp One incremented salary is: ", incrementedSalary)

	empTwo := EmpIncrementStruct{"Meha", 200, 50}
	incrementedSalaryTwo := empTwo.incrementedSalary()
	fmt.Println("Emp Two incremented salary is: ", incrementedSalaryTwo)
	fmt.Println(empOne.salary)
}
