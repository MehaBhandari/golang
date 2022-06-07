package main

import "fmt"

// One Type
type EmployeeStruct struct {
	empName string
	empAge byte
	empDesignation string
}

// Other Type
type EmployeeData struct {
	empBloodGroup string
	empSalaryBand byte
}

func main() {
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