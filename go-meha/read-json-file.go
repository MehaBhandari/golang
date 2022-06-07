package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type empJsonData struct {
	empJsonData [] empObj `json: "emp_obj"`
}

type empObj struct {
	userId string `json:"userId"`
    jobTitleName string `json:"jobTitleName"`
    firstName string `json:"firstName"`
    lastName string `json:"lastName"`
    preferredFullName string `json:"preferredFullName"`
    employeeCode string `json:"employeeCode"`
    region string `json:"region"`
    phoneNumber string `json:"phoneNumber`
    emailAddress string `json:"emailAddress`
}

func readJson() {
	file, _ := ioutil.ReadFile("employee-details.json")
	data := empJsonData{}
	_ = json.Unmarshal([]byte(file), &data)
	for i:=0; i < len(data.empJsonData); i++ {
		fmt.Println("User Id: ", data.empJsonData[i].userId)
		fmt.Println("Job Title: ", data.empJsonData[i].jobTitleName)
		fmt.Println("First Name: ", data.empJsonData[i].firstName)
		fmt.Println("Full Name: ", data.empJsonData[i].preferredFullName)
		fmt.Println("Employee Code: ", data.empJsonData[i].employeeCode)
		fmt.Println("Region: ", data.empJsonData[i].region)
		fmt.Println("Phone Number: ", data.empJsonData[i].phoneNumber)
		fmt.Println("Email Address: ", data.empJsonData[i].emailAddress)		
	}
}