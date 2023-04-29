package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateRender, _ := template.ParseFiles("./templates/" + tmpl)
	err := templateRender.Execute(w, nil)
	if err != nil {
		fmt.Println("Error al cargar el template:", err)
		return
	}
}
