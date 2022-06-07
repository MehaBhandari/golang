package main

import "fmt"

// Type One
type EmployeeStruct struct {
	empName string
	empAge byte
	empDesignation string
}

// Type Two
type EmployeeData struct {
	empBloodGroup string
	empSalaryBand byte
}

func conventionalStructs() {
	employeeOne := EmployeeStruct {
		empName: "Meha",
		empAge: 30,
		empDesignation: "Developer",
	}

	employeeTwo := new(EmployeeData)
	employeeTwo.empBloodGroup = "A+ve"
	employeeTwo.empSalaryBand = 2

	fmt.Println("Employee One is: ", employeeOne)
	fmt.Println("Employee Two is: ", *employeeTwo)
}