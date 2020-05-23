package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	templateString := `Lemonade Stand Supply`
	t := template.New("title")
	p, err := t.Parse(templateString)
	if err != nil {
		log.Println(err)
	}
	err = p.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}
}
