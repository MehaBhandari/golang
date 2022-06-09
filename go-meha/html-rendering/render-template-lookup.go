package main

import (
	"net/http"
)

var CheckPage = CheckPageStruct{isLookupPage: true}

func main() {			
	http.HandleFunc("/", AboutTechnofunnel)
	http.HandleFunc("/welcome-page", ShowHtmlTemplate)
	HandleScripts()
	http.ListenAndServe(":8000", nil)	
}


