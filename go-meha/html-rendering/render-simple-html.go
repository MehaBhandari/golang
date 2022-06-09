package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

type CheckPageStruct struct {
	isLookupPage bool
}

func HandleScripts() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("appScripts"))))
}

func renderSimpleHTML() {
	CheckPage = CheckPageStruct{isLookupPage: false}
	http.HandleFunc("/", ShowHtmlTemplate)
	http.HandleFunc("/about-Techno-Funnel", AboutTechnofunnel)
	HandleScripts()
	http.ListenAndServe(":8000", nil)
}

// localhost:8000/css/index.css

func ShowHtmlTemplate(response http.ResponseWriter, request *http.Request) {
	filePath := path.Join("html-pages","index.html")
	/* 
	In case of same folder
	filePath := path.Join("index.html")
	*/
	fmt.Println(filePath)
	applicationTemplate, _ := template.ParseFiles(filePath)
	applicationTemplate.Execute(response, CheckPage)
}


func AboutTechnofunnel(response http.ResponseWriter, request *http.Request) {
	filepath := path.Join("html-pages", "about.html")
	fmt.Println(filepath)
	applicationTemplate, _ :=template.ParseFiles(filepath)
	applicationTemplate.Execute(response, CheckPage)
}
