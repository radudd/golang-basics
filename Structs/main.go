package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var radu person
	radu.firstName = "Radu"
	radu.lastName = "D"
	radu.contact.email = "xxx@xxx.xxx"
	radu.contact.phone = "+123"
	
	// Pass by value, need pointers when updating structs
	//raduPointer := &radu
	//raduPointer.updateName("NewName")
	// works also like this - go does it for you
	radu.updateName("NewName")

	radu.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}
