package main

import (
	"fmt"
	"html/template"
	"os"
)
type TemplateStruct struct {
	Name string
	Age byte
}
func (temp TemplateStruct) ShowDetails() string{
	return "Name is " + temp.Name
}

func main() {
	newTemplate := TemplateStruct{"Meha", 30}

	fmt.Println("User name is: ", newTemplate.Name)

	templateString := `
		Hello {{ .Name}} & Age {{ .Age }}. 
		Welcome to GoLang Tutorials
		{{ .ShowDetails}}`
	customTemplate, _ := template.New("customTemplates").Parse(templateString)
	customTemplate.Execute(os.Stdout, newTemplate)
}