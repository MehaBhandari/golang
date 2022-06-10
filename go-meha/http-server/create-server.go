package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	Name string
	Age byte
}

func main() {
	http.HandleFunc("/", PrintSimpleString)
	http.HandleFunc("/showJsonData", PrintSimpleJson)
	http.ListenAndServe(":4000", nil)
}

func PrintSimpleString(w http.ResponseWriter, r *http.Request) {
	var userName = "Meha Bhandari"
	fmt.Fprintf(w , "Hello " +userName)
}

func PrintSimpleJson(w http.ResponseWriter, r *http.Request) {
	emp := Employee{"Gauri", 5}
	dataOutput, _ := json.Marshal(emp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataOutput)
}

// func ServerDataFiles(w http.ResponseWriter, r *http.Request) {
// 	http.ListenAndServe(":3000", http.FileServer(http.Dir("")))
// }