package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/radudd/golang-basics/ViewAdvanced/controller"
	"github.com/radudd/golang-basics/ViewAdvanced/viewModel"
)

func main() {
	templates := populateTemplates()
	controller.Startup
	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		context := viewModel.NewBase()
		if t != nil {
			err := t.Execute(w, context)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.ListenAndServe(":8080", nil)
}

// populate templates from folder
func populateTemplates() *template.Template {
	result := template.New("templates")
	const BasePath = "templates"
	template.Must(result.ParseGlob(BasePath + "/*.html"))
	return result
}
