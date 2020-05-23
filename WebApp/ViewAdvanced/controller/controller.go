package controller

import (
	"html/template"
	"net/http"
)

var homeController home

func Startup(template map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
