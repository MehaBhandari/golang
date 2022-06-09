package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func main() {
	http.HandleFunc("/", ShowHtmlTemplate)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("appScripts"))))
	http.ListenAndServe(":8000", nil)
}

// localhost:8000/css/index.css

func ShowHtmlTemplate(response http.ResponseWriter, request *http.Request) {
	filePath := path.Join("index.html")
	fmt.Println(filePath)
	applicationTemplate, _ := template.ParseFiles(filePath)
	applicationTemplate.Execute(response, nil)
}

