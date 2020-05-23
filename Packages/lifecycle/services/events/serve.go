// Package events provides a webservice that manage the library's events
package events

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/pluralsight/libmanager/services/internal/ports"
)

var port = 42

// StartServer register the handlesrs and initiate the webservice
func StartServer() error {
	sm := http.NewServeMux()
	sm.Handle("/", new(eventHandler))
	return http.ListenAndServe(":"+strconv.Itoa(port), sm)
}

func init() {
	//port := ports.EventService
	fmt.Println("serve.go 1", port)
	port := ports.EventService
	fmt.Println("serve.go 2", port)
}
