package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://cloud.redhat.com",
		"https://aws.amazon.com",
		"https://portal.azure.com",
	}

	c := make(chan string)
	//var c chan string

	for _, url := range urls {
		go callURL(url, c)
	}

	// iteration can be done also
	// by using the length of the slice
	/*
		for i := 0; i < len(urls); i++ {
			go callUrl(<-c)
		}
	*/

	/* infinite loop
	The go routine in this infine loops both:
	1. is a blocking call and
	   it consumes the information from the channel
	2. creates a new go routine for checking the URLs
	   with the url received in the channel
	*/

	/*
		for {
			go callUrl(<-c, c)
		}
	*/

	/*
		More common, a for loop using the range of channel
		can be used, i.e. wait for a value in the channel
		then assign it to l variable
	*/
	for u := range c {
		//go callUrl(l, c)
		/*
			but we will call this with a function literal (lambda)

			func() - defines the function literal
			() at the end invokes it

			We should never try to access a variable
			from main routine to a child routine

			As Go is pass by value language, we should provide
			the main routine variable as an argument for
			the literal function invoking the child routine
		*/
		go func(url string) {
			time.Sleep(2 * time.Second)
			callURL(url, c)
		}(u)
	}
}

func callURL(url string, c chan string) (bool, int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url + " is down :/")
		c <- url
		return false, (*resp).StatusCode
	}
	fmt.Println(url + " is UP !")
	c <- url
	return true, (*resp).StatusCode
}
