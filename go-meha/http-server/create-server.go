package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/simplePrint", PrintSimpleString)
}

func PrintSimpleString(w http.Response,r *http.Request) {
	
}