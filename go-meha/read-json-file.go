package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

type EmpObj struct {
	userId string 
    jobTitleName string 
    firstName string 
    lastName string 
    preferredFullName string 
    employeeCode string 
    region string 
    phoneNumber string 
    emailAddress string 
}

func readJsonFile() {
	file, err := ioutil.ReadFile("./employee-details.json")
	if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
	var payload []EmpObj
	err = json.Unmarshal(file, &payload)
	if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	log.Println("File is: ", string(file))
	log.Printf("origin: %s\n", payload)
	// for i:=0; i < len(data.empJsonData); i++ {
	// 	log.Println("User Id: ", payload.empJsonData[i].userId)
	// 	log.Println("Job Title: ", payload.empJsonData[i].jobTitleName)
	// 	log.Println("First Name: ", payload.empJsonData[i].firstName)
	// 	log.Println("Full Name: ", payload.empJsonData[i].preferredFullName)
	// 	log.Println("Employee Code: ", payload.empJsonData[i].employeeCode)
	// 	log.Println("Region: ", payload.empJsonData[i].region)
	// 	log.Println("Phone Number: ", payload.empJsonData[i].phoneNumber)
	// 	log.Println("Email Address: ", payload.empJsonData[i].emailAddress)		
	// }
}