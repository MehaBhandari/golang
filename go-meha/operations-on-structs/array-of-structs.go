package main

import "fmt"

func arrayOfStructs() {
	nameArray := [4]string {"A", "B", "C", "D"}
	ageArray := [4]byte {22, 25, 35, 45}
	bonusArray := [4]byte {10, 20, 30, 40}

	type EmployeeDataStruct struct {
		name string
		age byte
		bonus byte
	}

	var empSlice []EmployeeDataStruct
	

	for i:=0; i<len(nameArray); i++ {
		emp := new(EmployeeDataStruct)
		emp.name = nameArray[i]
		emp.age = ageArray[i]
		emp.bonus = bonusArray[i]
		empSlice = append(empSlice, *emp)
	}
	fmt.Println("empSlice is: ", empSlice)
}