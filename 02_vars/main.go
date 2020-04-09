package main

import "fmt"

func main() {
	// Using var
	//var name string = "Brad"

	//shorthand
	name, email := "Brad", "Brad@gmail.com"
	fmt.Println(name, email)
	fmt.Printf("%T\n", email)
}