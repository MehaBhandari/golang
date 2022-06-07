package main

import "fmt"

type EmployeeAddress struct {
	street, city, state, country string
	pincode int32
	isPermenant bool
}

type EmployeeDetails struct {
	name string
	age byte
	isMarried bool
	address EmployeeAddress
}

func main() {
	employeeOne := new(EmployeeDetails)
	employeeOne.name = "Meha"
	employeeOne.isMarried = true
	employeeOne.address.city = "New Delhi"
	employeeOne.address.isPermenant = true
	fmt.Println("Employee is: ", *employeeOne)
}