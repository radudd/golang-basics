package main

import (
	"fmt"
	"net/http"
)

type customHandler struct {
	greeting string
}

func (ch customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", ch.greeting)))
}

func main() {
	myHandler := customHandler{
		greeting: "Hello",
	}
	http.Handle("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
