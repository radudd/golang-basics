package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//fmt.Println(getText("http://localhost:8090/hello"))
	getText("http://localhost:8090/hello")
}

func getText(url string) *os.File {
	resp, errGet := http.Get(url)
	if errGet != nil {
		fmt.Println("Error by GET request:", errGet)
		os.Exit(1)
	}

	// instead of definining bs like before, i.e below
	// bs := []byte{}
	// do it with make, to initialize with n number of elements
	// we do this because Read can not increase the bs
	// and will add elements to bs until it is full, and with previous declaration it was empty
	// and read will see it as already full

	/*
		bs := make([]byte, 99999)
		noBytes, errRead := (*resp).Body.Read(bs)
		if errRead != nil {
			fmt.Println("Error by reading response:", errRead)
			os.Exit(1)
		}
		return noBytes, string(bs)
	*/

	// but easier to do it with io.Copy
	io.Copy(os.Stdout, resp.Body)
	return os.Stdout
}

// an element of the struct can can be of type interface, i.e must satisfy an interface
// i.e.c http.Response.Body -> io.ReadClose
// io.ReadClose is an interface which must satisfy other two interfaces
// Reader and Closer
// Reader is a common interfaces that is implemented by different kind of packages
// that will bring sources of data for your app, i.e. http.Response.Body
