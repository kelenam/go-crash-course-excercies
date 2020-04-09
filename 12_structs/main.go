package main

import (
	"fmt"
	"strconv"
)
// Define person struct
type Person struct {
	// firstName string 
	// lastName string 
	// city string
	// gender string
    // age int 
    firstName, lastName, city, gender string
    age  							  int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseName string) {
	p.lastName = spouseName
}

func main() {
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person1.hasBirthday()
	person1.getMarried("Smoley")
	fmt.Println(person1.greet())
}