package main

import "fmt"

type EmpStruct struct {
	name string
	designation string
	company string
}

func modifyEmployee(emp *EmpStruct) {
	emp.name = "Gauri"
	
}

func updatingStructs() {
	empOne := new(EmpStruct)
	empOne.name = "Meha"
	empOne.company = "GL"
	empOne.designation = "VP"
	fmt.Println("Emp One is: ", *empOne)
	//This won't modify the struct
	// modifyEmployee(*empOne) 
	modifyEmployee(empOne) //This will modify the struct
	fmt.Println("Modified Emp One is: ", *empOne)
	empTwo := EmpStruct{"Mayank", "Manager", "BlackRock"}
	fmt.Println("Emp Two is: ", empTwo)
}