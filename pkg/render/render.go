package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedFile, _ := template.ParseFiles("./templates" + tmpl)
	err := parsedFile.Execute(w, nil)

	log.Println("error occur while executing parsed file", err)

}
