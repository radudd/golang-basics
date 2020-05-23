package main

import (
	"log"
	"sync"
	"github.com/pluralsight/libmanager/services/events"
)

func main() {
	log.Println("Starting all services")

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go StartServer(wg, events.StartServer)
	wg.Wait()
	log.Println("All services stopped")
}

func startServer(wg *sync.WaitGroup, startFun func() err) {
	err := startFunc()
	wg.Done()
	if err != nil {
		log.Fatal(err)
	}
}
