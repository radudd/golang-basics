package main

import "fmt"

// interface for both bots
// it matches all the structures that has an associated 
// getGreeting which returns a string
// associated - receiver function

type bot interface {
	getGreeting() (string,error)
}

// interfaces are implicit
// you don't need to specify it implements in interfaces
// go takes care for you by checking the associated functions
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// receiver instance can be removed if not used
// the receiver can contain ONLY the receiver type
func (englishBot) getGreeting() (string,error) {
	//very custom login for english bot
	return "Hey there",nil
}

func (spanishBot) getGreeting() (string,error) {
	//very custom login for spanish bot
	return "Hola!",nil
}

// Go doesn't support overloading
/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb englishBot) {
	fmt.Println(sb.getGreeting())
}
*/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}